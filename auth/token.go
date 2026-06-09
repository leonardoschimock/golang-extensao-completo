package auth

import (
	"extensao-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userID
	claims["exp"] = time.Now().Add(6 * time.Hour).Unix()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	return token.SignedString(config.SecretKey)
}
