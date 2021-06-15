package controller

import (
	"github.com/gin-gonic/gin"
	"kada-account/model"
	"kada-account/token"
	"net/http"
)

func UpdateProfile() func(c *gin.Context) {
	return func(c *gin.Context) {
		userInfo := new(model.UserInfo)
		e := c.Bind(userInfo)
		if e != nil {
			c.JSON(http.StatusBadRequest, model.GetNetError("Profile verify error", e))
			c.Abort()
			return
		}
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		accessToken := authHeader[len(BearerSchema):]
		tokenClaims, tokenError := token.GetTokenClaimsFromToken(accessToken)
		if tokenError != nil {
			c.JSON(http.StatusInternalServerError, model.GetNetErrorWithCode(http.StatusInternalServerError, "Token parse fail", e))
			c.Abort()
			return
		}
		updateError := model.UpdateProfile(tokenClaims.UserId, userInfo)
		if updateError != nil {
			c.JSON(http.StatusBadRequest, model.GetNetErrorWithCode(http.StatusBadRequest, "Profile update fail", e))
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, userInfo)
	}
}
