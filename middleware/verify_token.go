package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/themethaithian/go-pos-system/app"
	"github.com/themethaithian/go-pos-system/config"
	"github.com/themethaithian/go-pos-system/token"
)

func (mdw *middleware) VerifyToken(next app.HandlerFunc) app.HandlerFunc {
	return func(ctx app.Context) {
		if config.Val.NeedAuthen {
			authHeader := ctx.Value("Authorization")
			if authHeader == "" {
				ctx.JSON(http.StatusUnauthorized, fmt.Errorf("authorization header not found"))
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				ctx.JSON(http.StatusUnauthorized, fmt.Errorf("invalid authorization header"))
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			_, err := token.VerifyToken(tokenString)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, errors.Wrapf(err, "failed to verify token"))
				return
			}
		}

		next(ctx)
	}
}
