package common

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePasswordHash(plainPassword, hasedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(plainPassword))
	return err == nil
}
