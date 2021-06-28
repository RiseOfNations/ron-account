package token

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ron-account/util"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, util.GetNetError("without token", nil))
			c.Abort()
			return
		}
		accessToken := authHeader[len(BearerSchema):]
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, util.GetNetError("auth error", nil))
			c.Abort()
			return
		} else {
			if VerifyToken(accessToken) == false {
				c.JSON(http.StatusUnauthorized, util.GetNetError("token error", nil))
				c.Abort()
				return
			}
		}
	}
}
