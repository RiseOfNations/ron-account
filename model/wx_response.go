package model

type WxResponse struct {
	OpenID     string `json:"openid" gorm:"type:varchar(50)"`
	SessionKey string `json:"session_key,omitempty" gorm:"type:varchar(50)"`
	UnionID    string `json:"unionid" gorm:"type:varchar(50)"`
}
