package product

import (
	"github.com/go-playground/validator/v10"
)

type Handler interface{}

type handler struct {
	validator *validator.Validate
	storage   Storage
}

func NewHandler(storage Storage, validator *validator.Validate) Handler {
	return &handler{
		storage:   storage,
		validator: validator,
	}
}
