package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHashPassword generates a hash from a password.
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CompareHashPassword compares a password (plain string) to a hash.
func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
