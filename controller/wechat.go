package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"kada-account/model"
	"kada-account/util"
	"net/http"
	"strings"
)

func WechatAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		code := c.Query("code")
		if strings.Compare(code, "") == 0 {
			c.JSON(http.StatusBadRequest, model.GetNetError("Without code.", nil))
			c.Abort()
			return
		}
		requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=wxadc83b40219cdbe5&secret=06a72d53ce36e6923566f803bbb3451e&js_code=" + code + "&grant_type=authorization_code"
		wxResp, wxError := http.Get(requestUrl)
		if wxError != nil {
			c.JSON(http.StatusInternalServerError, model.GetNetError("Get wx login info error.", wxError))
			c.Abort()
			return
		}
		defer wxResp.Body.Close()
		body, readWxBodyError := ioutil.ReadAll(wxResp.Body)
		if readWxBodyError != nil {
			c.JSON(http.StatusInternalServerError, model.GetNetError("Read wx body error.", readWxBodyError))
			c.Abort()
			return
		}
		var raw map[string]interface{}
		wxJsonError := util.JsonUnmarshal(body, &raw)
		if wxJsonError != nil {
			c.JSON(http.StatusInternalServerError, model.GetNetError("Read wx body error.", readWxBodyError))
			c.Abort()
			return
		}
		if raw["errcode"] != nil {
			c.JSON(http.StatusBadRequest, model.GetNetErrorWithCode(raw["errcode"].(int), raw["errmsg"].(string), nil))
			c.Abort()
			return
		} else {
			wxRespModel := new(model.WxResponse)
			_ = util.JsonUnmarshal(body, wxRespModel)
			user, exist := model.GetUserInfoByOpenId(wxRespModel.OpenID)
			if exist {
				user.SessionKey = wxRespModel.SessionKey
				_ = user.UpdateWxRespUserInfo()
				token, e := util.GenerateToken(user)
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
				user.OpenID = wxRespModel.OpenID
				user.UnionID = wxRespModel.UnionID
				user.SessionKey = wxRespModel.SessionKey
				_ = user.CreateUser()
				token, e := util.GenerateToken(user)
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
}
