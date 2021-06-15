package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"kada-account/model"
	token2 "kada-account/token"
	"net/http"
	"strings"
)

func SmsAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		smsCode := c.Query("sms_code")
		if strings.Compare(smsCode, "") == 0 {
			c.JSON(http.StatusBadRequest, model.GetNetError("Without sms code.", nil))
			c.Abort()
			return
		}
		phoneNumber := c.Query("phone_number")
		if strings.Compare(smsCode, "") == 0 {
			c.JSON(http.StatusBadRequest, model.GetNetError("Without phone number.", nil))
			c.Abort()
			return
		}
		// 添加教研逻辑
		if checkSmsCode(phoneNumber, smsCode) != nil {
			c.JSON(http.StatusBadRequest, model.GetNetErrorWithCode(1000, "sms code error.", nil))
			c.Abort()
			return
		}
		user, exist := model.GetUserInfoByPhoneNumber(phoneNumber)
		if exist {
			token, e := token2.GenerateToken(user)
			if e != nil {
				c.JSON(http.StatusInternalServerError, model.GetNetErrorWithCode(http.StatusInternalServerError, "Token generate fail", e))
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, &model.LoginResponse{
				Token: token,
			})
		} else {
			user := new(model.User)
			user.UserId = uuid.NewString()
			user.PhoneNumber = phoneNumber
			_ = user.CreateUser()
			token, e := token2.GenerateToken(user)
			if e != nil {
				c.JSON(http.StatusInternalServerError, model.GetNetErrorWithCode(http.StatusInternalServerError, "Token generate fail", e))
				c.Abort()
				return
			}
			c.JSON(http.StatusCreated, &model.LoginResponse{
				Token: token,
			})
		}
	}
}

func checkSmsCode(phoneNumber string, smsCode string) error {
	if smsCode == "1234" {
		return nil
	} else {
		return errors.New("sms code error")
	}
}
