package user

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/themethaithian/go-pos-system/app"
)

type NewUser struct {
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	Role            string `json:"role"`
}

func (u *NewUser) Validate() error {
	if !slices.Contains(app.VALIDATE_ROW, u.Role) {
		return fmt.Errorf("invalid role (got: %s)", u.Role)
	}

	if len(u.Password) < 8 {
		return fmt.Errorf("invalid password (minimum length is 8)")
	}

	if u.Password != u.ConfirmPassword {
		return fmt.Errorf("invalid password (password and confirm is not the same)")
	}

	return nil
}

func (h *handler) CreateUser(ctx app.Context) {
	var request NewUser
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

	user := mappingUser(request)

	if err := h.storage.InsertNewUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func mappingUser(request NewUser) User {
	return User{
		Username:     request.Username,
		PasswordHash: "",
		Role:         request.Role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
