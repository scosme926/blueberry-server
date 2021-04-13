package utils

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(w http.ResponseWriter, r *http.Request, passkey string) []byte {
	password := []byte(passkey)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return hashedPassword
}

func CompareHashedPassword(w http.ResponseWriter, r *http.Request, hashedpasskey []byte, requestPasskey []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedpasskey, requestPasskey)
	if err != nil {
		return false
	}
	return true
}