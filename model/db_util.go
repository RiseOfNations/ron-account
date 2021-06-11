package model

import (
	"github.com/gin-gonic/gin"
	sqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GetDb() (*gorm.DB, error) {
	config := sqldriver.NewConfig()
	config.User = "kada"
	config.Passwd = "zQaRxNcF9yH1NBgY"
	if gin.Mode() == gin.DebugMode {
		config.Addr = "127.0.0.1:3306"
	} else {
		config.Addr = "rm-hp3f2888fmegv2ktc.mysql.huhehaote.rds.aliyuncs.com"
	}
	config.DBName = "kada"
	config.ParseTime = true
	config.Loc = time.UTC
	config.Net = "tcp"
	config.Params = map[string]string{
		"charset": "utf8mb4",
	}
	return gorm.Open(mysql.Open(config.FormatDSN()), &gorm.Config{})
}

func InstDb() bool {
	db, errConnDb := GetDb()
	if errConnDb != nil {
		println(errConnDb)
		return false
	}
	sqlDb, connErr := db.DB()
	if connErr != nil {
		println(connErr.Error())
		return false
	}
	pingError := sqlDb.Ping()
	if pingError != nil {
		println(pingError.Error())
		return false
	}
	err := db.AutoMigrate(&User{})
	if err != nil {
		print(err)
		return false
	}
	return true
}
