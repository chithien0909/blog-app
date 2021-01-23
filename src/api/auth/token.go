package auth

import (
	"../../config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(userId uint64) (string, error)  {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.SECRET_KEY)
}