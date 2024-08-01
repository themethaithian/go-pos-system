package app

import (
	"encoding/json"
	"net/http"
)

type contextHTTP struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *contextHTTP) Bind(v any) error {
	defer c.r.Body.Close()
	return json.NewDecoder(c.r.Body).Decode(&v)
}

func (c *contextHTTP) Param(key string) string {
	return c.r.PathValue(key)
}

func (c *contextHTTP) JSON(statusCode int, v any) {
	c.w.WriteHeader(statusCode)

	switch v := v.(type) {
	case error:
		jsonErr := json.NewEncoder(c.w).Encode(Response{
			Status:  "SUCCESS",
			Message: v.Error(),
		})
		_ = jsonErr

		return
	}

	err := json.NewEncoder(c.w).Encode(Response{
		Status: "FAILED",
		Data:   v,
	})
	_ = err
}

type RouterHTTP struct {
	mux          *http.ServeMux
	interceptors []middlewareFunc
}

func NewRouterHTTP() *RouterHTTP {
	return &RouterHTTP{mux: http.NewServeMux()}
}

type middlewareFunc func(h http.Handler) http.Handler

func (router *RouterHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

func (router *RouterHTTP) Use(h ...middlewareFunc) {
	router.interceptors = append(router.interceptors, h...)
}

func (router *RouterHTTP) GET(path string, handlerFn HandlerFunc) {
	router.mux.Handle(path, NewHTTPHandler(http.MethodGet, handlerFn, router.interceptors))
}

func (router *RouterHTTP) POST(path string, handlerFn HandlerFunc) {
	router.mux.Handle(path, NewHTTPHandler(http.MethodPost, handlerFn, router.interceptors))
}

func NewHTTPHandler(method string, handler func(Context), interceptors []middlewareFunc) http.Handler {
	var httpHandler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(&contextHTTP{w: w, r: r})
	})

	for _, interceptor := range interceptors {
		httpHandler = interceptor(http.Handler(httpHandler))
	}
	return httpHandler
}
