package geolab_core

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	return string(hash), err
}
