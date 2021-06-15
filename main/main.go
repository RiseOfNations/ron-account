package main

import (
	"github.com/gin-gonic/gin"
	"kada-account/controller"
	"kada-account/model"
	"kada-account/token"
)

func main() {
	gin.SetMode(gin.DebugMode)
	if !model.InstDb() {
		println("inst db error.")
		return
	}
	router := gin.Default()
	authorRouter := router.Group("/auth")
	{
		authorRouter.GET("/wechat", controller.WechatAuth())
		authorRouter.GET("/sms", controller.SmsAuth())
	}
	userRouter := router.Group("/user")
	{
		userRouter.PUT("/profile", token.Middleware(), controller.UpdateProfile())
	}
	_ = router.Run(":10002")
}
