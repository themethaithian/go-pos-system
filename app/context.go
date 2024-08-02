package app

type Context interface {
	Bind(v any) error
	Param(key string) string
	JSON(statusCode int, v any)
	PathValue(v string) string
}

type HandlerFunc func(Context)
