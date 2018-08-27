package customer

import (
	"golang.org/x/crypto/bcrypt"
)

type Password string

const LENGTH = 10

func (p Password) Hash() (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), LENGTH)
	return string(hashed), err
}

func (p Password) ComparePassword(hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
}
