package login

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	token2 "ron-account/token"
	user2 "ron-account/user"
	"ron-account/util"
	"strings"
)

func WechatAuthController() func(c *gin.Context) {
	return func(c *gin.Context) {
		code := c.Query("code")
		if strings.Compare(code, "") == 0 {
			c.JSON(http.StatusBadRequest, util.GetNetError("Without code.", nil))
			c.Abort()
			return
		}
		requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=wxadc83b40219cdbe5&secret=06a72d53ce36e6923566f803bbb3451e&js_code=" + code + "&grant_type=authorization_code"
		wxResp, wxError := http.Get(requestUrl)
		if wxError != nil {
			c.JSON(http.StatusInternalServerError, util.GetNetError("Get wx login info net.", wxError))
			c.Abort()
			return
		}
		defer wxResp.Body.Close()
		body, readWxBodyError := ioutil.ReadAll(wxResp.Body)
		if readWxBodyError != nil {
			c.JSON(http.StatusInternalServerError, util.GetNetError("Read wx body net.", readWxBodyError))
			c.Abort()
			return
		}
		var raw map[string]interface{}
		wxJsonError := util.Unmarshal(body, &raw)
		if wxJsonError != nil {
			c.JSON(http.StatusInternalServerError, util.GetNetError("Read wx body net.", readWxBodyError))
			c.Abort()
			return
		}
		if raw["errcode"] != nil {
			c.JSON(http.StatusBadRequest, util.GetNetErrorWithCode(raw["errcode"].(int), raw["errmsg"].(string), nil))
			c.Abort()
			return
		} else {
			wxRespModel := new(user2.WxResponse)
			_ = util.Unmarshal(body, wxRespModel)
			user, exist := user2.GetUserInfoByOpenId(wxRespModel.OpenID)
			if exist {
				user.SessionKey = wxRespModel.SessionKey
				_ = user.UpdateWxRespUserInfo()
				token, e := token2.GenerateToken(user.UserId, user.Nickname, user.AvatarUrl)
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
				user.OpenID = wxRespModel.OpenID
				user.UnionID = wxRespModel.UnionID
				user.SessionKey = wxRespModel.SessionKey
				_ = user.CreateUser()
				token, e := token2.GenerateToken(user.UserId, user.Nickname, user.AvatarUrl)
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
}
