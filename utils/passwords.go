package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword hashes the given password using bcrypt
func EncryptPassword(password string) (string, error) {
	// Generate a salted hash for the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Convert the hashed password to a string and return
	return string(hashedPassword), nil
}

// VerifyPassword checks if the provided password matches the stored hash
func VerifyPassword(inputPassword, storedHash string) bool {
	// Compare the input password with the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(inputPassword))
	return err == nil
}
