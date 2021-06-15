package token

import (
	"github.com/gin-gonic/gin"
	"kada-account/model"
	"net/http"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		accessToken := authHeader[len(BearerSchema):]
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, model.GetNetError("auth error", nil))
			c.Abort()
		} else {
			if VerifyToken(accessToken) == false {
				c.JSON(http.StatusUnauthorized, model.GetNetError("token error", nil))
				c.Abort()
			}
		}
	}
}
