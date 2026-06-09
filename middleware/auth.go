package middleware

import (
	"extensao-api/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := extractToken(r)

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return config.SecretKey, nil
			},
		)

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func extractToken(r *http.Request) string {

	bearer := r.Header.Get("Authorization")

	if bearer == "" {
		return ""
	}

	parts := strings.Split(bearer, " ")

	if len(parts) != 2 {
		return ""
	}

	return parts[1]
}
