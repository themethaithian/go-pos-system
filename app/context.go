package app

type Context interface {
	Bind(v any) error
	Param(key string) string
	JSON(statusCode int, v any)
}

type HandlerFunc func(Context)
