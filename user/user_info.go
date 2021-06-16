package user

type UserInfo struct {
	NickName  string `util:"nick_name,omitempty" gorm:"type:varchar(20)"`
	AvatarUrl string `util:"avatar_url,omitempty" gorm:"type:varchar(255)"`
	Gender    string `util:"gender,omitempty" gorm:"type:varchar(10)"`
	City      string `util:"city,omitempty" gorm:"type:varchar(50)"`
	Province  string `util:"province,omitempty" gorm:"type:varchar(50)"`
	Country   string `util:"country,omitempty" gorm:"type:varchar(20)"`
	Language  string `util:"language,omitempty" gorm:"type:varchar(10)"`
}
