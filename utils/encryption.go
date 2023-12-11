package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return bytes, err
}

func CompareHashedPasswords(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
