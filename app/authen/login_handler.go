package authen

import (
	"fmt"
	"net/http"

	"github.com/themethaithian/go-pos-system/app"
	"github.com/themethaithian/go-pos-system/token"
)

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *handler) Login(ctx app.Context) {
	var login Login

	if err := ctx.Bind(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("incorrect username or password"))
		return
	}

	if err := h.validator.Struct(login); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	role := "ADMIN" // TODO: get role from tbl_user

	var res LoginResponse

	accessToken, err := token.CreateJWT(login.Username, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	res.AccessToken = accessToken

	ctx.JSON(http.StatusOK, res)
}
