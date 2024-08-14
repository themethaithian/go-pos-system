package product

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/themethaithian/go-pos-system/app"
)

type NewProduct struct {
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

	if err := h.storage.InsertProduct(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
