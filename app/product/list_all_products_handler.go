package product

import (
	"net/http"

	"github.com/themethaithian/go-pos-system/app"
)

func (h *handler) ListAllProducts(ctx app.Context) {
	products, err := h.storage.RetrieveAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}
