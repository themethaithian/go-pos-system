package authen

import (
	"fmt"
	"net/http"

	"github.com/themethaithian/go-pos-system/app"
)

type Register struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Password2 string `json:"password2" validate:"required"`
}

func (h *handler) Register(ctx app.Context) {
	var register Register

	if err := ctx.Bind(&register); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(register); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if register.Password != register.Password2 {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid password"))
		return
	}

	err := h.storage.CreateUser(register)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
