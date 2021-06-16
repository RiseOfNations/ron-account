package token

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	UserId string `util:"user_id"`
}
