package product

import (
	"net/http"

	"github.com/themethaithian/go-pos-system/app"
)

type EditProduct struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float32 `json:"price"`
	Quantity    *int     `json:"quantity"`
}

func (h *handler) EditProduct(ctx app.Context) {
	var editProduct EditProduct
	if err := ctx.Bind(&editProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id := ctx.PathValue("id")

	err := h.storage.UpdateProduct(id, editProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
