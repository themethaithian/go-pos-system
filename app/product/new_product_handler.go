package product

import (
	"net/http"

	"github.com/themethaithian/go-pos-system/app"
)

type NewProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

func (h *handler) NewProduct(ctx app.Context) {
	var newProductReq NewProductRequest
	if err := ctx.Bind(&newProductReq); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(newProductReq); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
