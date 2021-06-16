package token

import (
	"github.com/gin-gonic/gin"
	"kada-account/util"
	"net/http"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		accessToken := authHeader[len(BearerSchema):]
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, util.GetNetError("auth net", nil))
			c.Abort()
		} else {
			if VerifyToken(accessToken) == false {
				c.JSON(http.StatusUnauthorized, util.GetNetError("token net", nil))
				c.Abort()
			}
		}
	}
}
