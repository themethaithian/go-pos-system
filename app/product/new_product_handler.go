package product

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/themethaithian/go-pos-system/app"
)

type NewProduct struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

func (h *handler) NewProduct(ctx app.Context) {
	var newProductReq NewProduct
	if err := ctx.Bind(&newProductReq); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(newProductReq); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id := uuid.New().String()

	if err := h.storage.InsertProduct(id, newProductReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
