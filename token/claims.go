package token

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Nickname  string `json:"nickname,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	UserId    string `json:"user_id"`
}
