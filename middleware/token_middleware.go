package middleware

import (
	"github.com/gin-gonic/gin"
	"kada-account/model"
	"kada-account/util"
	"net/http"
)

func TokenMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, model.GetNetError("auth error", nil))
			c.Abort()
		} else {
			if util.VerifyToken(accessToken) == false {
				c.JSON(http.StatusUnauthorized, model.GetNetError("token error", nil))
				c.Abort()
			}
		}
	}
}
