package db

import (
	"github.com/gin-gonic/gin"
	sqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GetDb() (*gorm.DB, error) {
	config := sqldriver.NewConfig()
	config.User = "ron"
	config.Passwd = "zQaRxNcF9yH1NBgY"
	if gin.Mode() == gin.DebugMode {
		config.Addr = "127.0.0.1:3306"
	} else {
		config.Addr = "rm-2zeyd1yxsq3e4mycb.mysql.rds.aliyuncs.com:3306"
	}
	config.DBName = "ron"
	config.ParseTime = true
	config.Loc = time.UTC
	config.Net = "tcp"
	config.Params = map[string]string{
		"charset": "utf8mb4",
	}
	return gorm.Open(mysql.Open(config.FormatDSN()), &gorm.Config{})
}
