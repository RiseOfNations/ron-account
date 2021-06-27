package token

import (
	"github.com/dgrijalva/jwt-go"
	"ron-account/user"
)

type Claims struct {
	jwt.StandardClaims
	user.Info
}
