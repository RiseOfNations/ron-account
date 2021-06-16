package login

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	token2 "kada-account/token"
	user2 "kada-account/user"
	"kada-account/util"
	"net/http"
	"strings"
)

func SmsAuthController() func(c *gin.Context) {
	return func(c *gin.Context) {
		smsCode := c.Query("sms_code")
		if strings.Compare(smsCode, "") == 0 {
			c.JSON(http.StatusBadRequest, util.GetNetError("Without sms code.", nil))
			c.Abort()
			return
		}
		phoneNumber := c.Query("phone_number")
		if strings.Compare(smsCode, "") == 0 {
			c.JSON(http.StatusBadRequest, util.GetNetError("Without phone number.", nil))
			c.Abort()
			return
		}
		// 添加教研逻辑
		if checkSmsCode(phoneNumber, smsCode) != nil {
			c.JSON(http.StatusBadRequest, util.GetNetErrorWithCode(1000, "sms code net.", nil))
			c.Abort()
			return
		}
		user, exist := user2.GetUserInfoByPhoneNumber(phoneNumber)
		if exist {
			token, e := token2.GenerateToken(user.UserId)
			if e != nil {
				c.JSON(http.StatusInternalServerError, util.GetNetErrorWithCode(http.StatusInternalServerError, "Token generate fail", e))
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, &Response{
				Token: token,
			})
		} else {
			user := new(user2.User)
			user.UserId = uuid.NewString()
			user.PhoneNumber = phoneNumber
			_ = user.CreateUser()
			token, e := token2.GenerateToken(user.UserId)
			if e != nil {
				c.JSON(http.StatusInternalServerError, util.GetNetErrorWithCode(http.StatusInternalServerError, "Token generate fail", e))
				c.Abort()
				return
			}
			c.JSON(http.StatusCreated, &Response{
				Token: token,
			})
		}
	}
}

func checkSmsCode(phoneNumber string, smsCode string) error {
	if smsCode == "1234" {
		return nil
	} else {
		return errors.New("sms code net")
	}
}