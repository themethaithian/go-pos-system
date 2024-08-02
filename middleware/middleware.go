package middleware

import "net/http"

type Middleware interface {
	VerifyToken(handler http.Handler) http.Handler
}

type middleware struct{}

func New() Middleware {
	return &middleware{}
}
