package user

import (
	"github.com/go-playground/validator/v10"

	"github.com/themethaithian/go-pos-system/app"
	"github.com/themethaithian/go-pos-system/hashing"
)

type Handler interface {
	CreateUser(ctx app.Context)
}

type handler struct {
	validator validator.Validate
	hashing   hashing.Hashing
	storage   Storage
}

func NewHandler(storage Storage) Handler {
	return &handler{
		storage: storage,
	}
}
