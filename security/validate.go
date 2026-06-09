package security

import "golang.org/x/crypto/bcrypt"

func ValidatePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
