package middleware

import (
	"github.com/themethaithian/go-pos-system/app"
)

type Middleware interface {
	VerifyToken(next app.HandlerFunc) app.HandlerFunc
}

type middleware struct{}

func New() Middleware {
	return &middleware{}
}
