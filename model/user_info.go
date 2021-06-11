package model

type UserInfo struct {
	NickName  string `json:"nick_name" gorm:"type:varchar(20)"`
	AvatarUrl string `json:"avatar_url" gorm:"type:varchar(255)"`
	Gender    string `json:"gender" gorm:"type:varchar(10)"`
	City      string `json:"city" gorm:"type:varchar(50)"`
	Province  string `json:"province" gorm:"type:varchar(50)"`
	Country   string `json:"country" gorm:"type:varchar(20)"`
	Language  string `json:"language" gorm:"type:varchar(10)"`
}
