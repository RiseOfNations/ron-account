package user

type WxResponse struct {
	OpenID     string `util:"openid" gorm:"type:varchar(50)"`
	SessionKey string `util:"session_key,omitempty" gorm:"type:varchar(50)"`
	UnionID    string `util:"unionid" gorm:"type:varchar(50)"`
}
