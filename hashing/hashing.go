package hashing

import "golang.org/x/crypto/bcrypt"

type Hashing interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
}

type hashing struct{}

func NewHashing() Hashing {
	return &hashing{}
}

func (h *hashing) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *hashing) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
