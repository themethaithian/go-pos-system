package product

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/themethaithian/go-pos-system/app"
)

type NewProduct struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Price       float32 `json:"price" validate:"required,gte=0"`
}

func (h *handler) NewProduct(ctx app.Context) {
	var newProduct NewProduct
	if err := ctx.Bind(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id := uuid.New().String()

	if err := h.storage.InsertProduct(id, newProduct); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
