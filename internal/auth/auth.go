package auth

import (
	"github.com/alexedwards/argon2id"
)

// Validate the email format
// Check the email doesn't already exist in the DB
// Hash the password with Argon2id
// Insert the user into the DB
// Return the created user

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func CheckHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return match, nil
}
