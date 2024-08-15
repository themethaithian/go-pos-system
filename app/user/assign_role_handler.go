package user

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/themethaithian/go-pos-system/app"
)

type AssignRole struct {
	Id   int    `json:"id" validate:"required"`
	Role string `json:"role" validate:"required"`
}

func (u *AssignRole) Validate() error {
	if !slices.Contains(app.VALIDATE_ROW, u.Role) {
		return fmt.Errorf("invalid role (got: %s)", u.Role)
	}

	return nil
}

func (h *handler) AssignRole(ctx app.Context) {
	var request AssignRole

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.storage.UpdateUserRole(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
