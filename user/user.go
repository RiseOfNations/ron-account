package user

import (
	"errors"
	"gorm.io/gorm"
	"ron-account/db"
	"ron-account/util"
	"time"
)

type User struct {
	UserInfo
	WxResponse
	UserId      string    `json:"user_id" gorm:"primaryKey;type:varchar(36);not null"`
	PhoneNumber string    `json:"-" gorm:"type:varchar(16)"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// GetUserInfoByOpenId 通过微信open id获取用户
func GetUserInfoByOpenId(openID string) (*User, bool) {
	db, openSqlError := db.GetDb()
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

// GetUserInfoByPhoneNumber 通过微信手机号获取用户
func GetUserInfoByPhoneNumber(phoneNumber string) (*User, bool) {
	db, openSqlError := db.GetDb()
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

// GetUserInfoByUserId 通过user id获取用户
func GetUserInfoByUserId(userId string) (*User, bool) {
	db, openSqlError := db.GetDb()
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

func UpdateProfile(userId string, userInfo *UserInfo) error {
	// 添加检查逻辑
	db, openSqlError := db.GetDb()
	if openSqlError != nil {
		return openSqlError
	}
	// 转换成map
	data, marshalError := util.Marshal(userInfo)
	if marshalError != nil {
		return errors.New("user info verify fail")
	}
	updateMap := map[string]interface{}{}
	unmarshalError := util.Unmarshal(data, &updateMap)
	if unmarshalError != nil {
		return errors.New("user info verify fail")
	}
	db.Model(&User{}).Where("user_id = ?", userId).Updates(updateMap)
	return nil
}

func (user *User) UpdateWechatUser() error {
	db, openSqlError := db.GetDb()
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
	db, openSqlError := db.GetDb()
	if openSqlError != nil {
		return openSqlError
	}
	return db.Model(user).Updates(map[string]interface{}{
		"session_key": user.SessionKey,
	}).Error
}

func (user *User) CreateUser() error {
	db, openSqlError := db.GetDb()
	if openSqlError != nil {
		return openSqlError
	}
	return db.Create(user).Error
}
