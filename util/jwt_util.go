package util

import (
	"github.com/dgrijalva/jwt-go"
	"kada-account/model"
)

var secretKey = "Dx5jdNFidWjL62eRqtC7Q"

// GenerateToken 没做刷新，需要补充
func GenerateToken(user *model.User) (string, error) {
	tokenClams := new(model.TokenClaims)
	tokenClams.StandardClaims = jwt.StandardClaims{
		ExpiresAt: jwt.TimeFunc().Unix(),
		Issuer:    "kada-account",
		IssuedAt:  jwt.TimeFunc().Unix(),
	}
	tokenClams.UserId = user.UserId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClams)
	return token.SignedString([]byte(secretKey))
}

// VerifyToken 验证token是否正常
func VerifyToken(token string) bool {
	tokenClaims := new(model.TokenClaims)
	rawToken, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false
	}
	if rawToken != nil {
		return rawToken.Valid
	} else {
		return false
	}
}

// GetTokenClaimsFromToken 获取auth信息
func GetTokenClaimsFromToken(token string) (*model.TokenClaims, error) {
	tokenClaims := new(model.TokenClaims)
	rawToken, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if rawToken == nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return rawToken.Claims.(*model.TokenClaims), nil
}
