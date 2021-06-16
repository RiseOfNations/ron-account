package main

import (
	"github.com/gin-gonic/gin"
	"kada-account/db"
	"kada-account/login"
	"kada-account/token"
	"kada-account/user"
)

func main() {
	gin.SetMode(gin.DebugMode)
	if !instDb() {
		println("inst db net.")
		return
	}
	router := gin.Default()
	authorRouter := router.Group("/auth")
	{
		authorRouter.GET("/wechat", login.WechatAuthController())
		authorRouter.GET("/sms", login.SmsAuthController())
	}
	userRouter := router.Group("/user")
	{
		userRouter.PUT("/profile", token.Middleware(), user.UpdateProfileController())
	}
	_ = router.Run(":10002")
}

func instDb() bool {
	getDb, errConnDb := db.GetDb()
	if errConnDb != nil {
		println(errConnDb)
		return false
	}
	sqlDb, connErr := getDb.DB()
	if connErr != nil {
		println(connErr.Error())
		return false
	}
	pingError := sqlDb.Ping()
	if pingError != nil {
		println(pingError.Error())
		return false
	}
	err := getDb.AutoMigrate(&user.User{})
	if err != nil {
		print(err)
		return false
	}
	return true
}
