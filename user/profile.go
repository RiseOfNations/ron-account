package user

import (
	"github.com/gin-gonic/gin"
	"kada-account/token"
	"kada-account/util"
	"net/http"
)

func UpdateProfileController() func(c *gin.Context) {
	return func(c *gin.Context) {
		userInfo := new(UserInfo)
		e := c.Bind(userInfo)
		if e != nil {
			c.JSON(http.StatusBadRequest, util.GetNetError("Profile verify net", e))
			c.Abort()
			return
		}
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		accessToken := authHeader[len(BearerSchema):]
		tokenClaims, tokenError := token.GetTokenClaimsFromToken(accessToken)
		if tokenError != nil {
			c.JSON(http.StatusInternalServerError, util.GetNetErrorWithCode(http.StatusInternalServerError, "Token parse fail", e))
			c.Abort()
			return
		}
		updateError := UpdateProfile(tokenClaims.UserId, userInfo)
		if updateError != nil {
			c.JSON(http.StatusBadRequest, util.GetNetErrorWithCode(http.StatusBadRequest, "Profile update fail", e))
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, userInfo)
	}
}