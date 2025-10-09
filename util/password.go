package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash passowrd %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassowrd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassowrd), []byte(password))
}
