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
	var request Login

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("incorrect username or password"))
		return
	}

	if err := h.validator.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	role, err := h.storage.RetrieveRoleFromUsername(request.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var response LoginResponse

	accessToken, err := token.CreateJWT(request.Username, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response.AccessToken = accessToken

	ctx.JSON(http.StatusOK, response)
}
