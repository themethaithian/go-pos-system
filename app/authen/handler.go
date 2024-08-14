package authen

import (
	"github.com/go-playground/validator/v10"

	"github.com/themethaithian/go-pos-system/app"
)

type Handler interface {
	Login(ctx app.Context)
}

type handler struct {
	validator *validator.Validate
	storage   Storage
}

func NewHandler(validator *validator.Validate, storage Storage) Handler {
	return &handler{
		validator: validator,
		storage:   storage,
	}
}
