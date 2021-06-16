package token

import (
	"github.com/dgrijalva/jwt-go"
)

var secretKey = "Dx5jdNFidWjL62eRqtC7Q"

// GenerateToken 没做刷新，需要补充
func GenerateToken(userId string) (string, error) {
	tokenClams := new(Claims)
	tokenClams.StandardClaims = jwt.StandardClaims{
		Issuer:    "kada-account",
		IssuedAt:  jwt.TimeFunc().Unix(),
	}
	tokenClams.UserId = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClams)
	return token.SignedString([]byte(secretKey))
}

// VerifyToken 验证token是否正常
func VerifyToken(token string) bool {
	tokenClaims := new(Claims)
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
func GetTokenClaimsFromToken(token string) (*Claims, error) {
	tokenClaims := new(Claims)
	rawToken, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if rawToken == nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return rawToken.Claims.(*Claims), nil
}
