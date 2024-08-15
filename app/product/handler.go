package product

import (
	"github.com/go-playground/validator/v10"

	"github.com/themethaithian/go-pos-system/app"
)

type Handler interface {
	ListAllProducts(app.Context)
	CreateProduct(app.Context)
	UpdateProduct(app.Context)
	DeleteProduct(app.Context)
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
