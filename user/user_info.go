package user

type Info struct {
	Nickname  string `json:"nickname,omitempty" gorm:"type:varchar(20)"`
	AvatarUrl string `json:"avatar_url,omitempty" gorm:"type:varchar(255)"`
	Gender    string `json:"gender,omitempty" gorm:"type:varchar(10)"`
	City      string `json:"city,omitempty" gorm:"type:varchar(50)"`
	Province  string `json:"province,omitempty" gorm:"type:varchar(50)"`
	Country   string `json:"country,omitempty" gorm:"type:varchar(20)"`
	Language  string `json:"language,omitempty" gorm:"type:varchar(10)"`
}
