package token

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}
