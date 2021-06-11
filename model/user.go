package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserInfo
	WxResponse
	UserId      string    `json:"kada_user_id" gorm:"primaryKey;type:varchar(36);not null"`
	PhoneNumber string    `json:"-" gorm:"type:varchar(16)"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func GetUserInfoByOpenId(openID string) (*User, bool) {
	db, openSqlError := GetDb()
	if openSqlError != nil {
		return nil, false
	}
	user := new(User)
	e := db.Where("open_id = ?", openID).First(user).Error
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, false
	}
	return user, true
}

func GetUserInfoByPhoneNumber(phoneNumber string) (*User, bool) {
	db, openSqlError := GetDb()
	if openSqlError != nil {
		return nil, false
	}
	user := new(User)
	e := db.Where("phone_number = ?", phoneNumber).First(user).Error
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, false
	}
	return user, true
}

func GetUserInfoByUserId(userId string) (*User, bool) {
	db, openSqlError := GetDb()
	if openSqlError != nil {
		return nil, false
	}
	user := new(User)
	e := db.Where("user_id = ?", userId).First(user).Error
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, false
	}
	return user, true
}

func (user *User) UpdateWechatUser() error {
	db, openSqlError := GetDb()
	if openSqlError != nil {
		return openSqlError
	}
	err := db.Model(user).Updates(map[string]interface{}{
		"open_id":    user.OpenID,
		"union_id":   user.UnionID,
		"nick_name":  user.NickName,
		"avatar_url": user.AvatarUrl,
		"gender":     user.Gender,
		"city":       user.City,
		"province":   user.Province,
		"country":    user.Country,
		"language":   user.Language,
	}).Error
	return err
}

func (user *User) UpdateWxRespUserInfo() error {
	db, openSqlError := GetDb()
	if openSqlError != nil {
		return openSqlError
	}
	return db.Model(user).Updates(map[string]interface{}{
		"session_key": user.SessionKey,
	}).Error
}

func (user *User) CreateUser() error {
	db, openSqlError := GetDb()
	if openSqlError != nil {
		return openSqlError
	}
	return db.Create(user).Error
}
