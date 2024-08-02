package product

import (
	"github.com/go-playground/validator/v10"

	"github.com/themethaithian/go-pos-system/app"
)

type Handler interface {
	NewProduct(ctx app.Context)
	EditProduct(ctx app.Context)
}

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
