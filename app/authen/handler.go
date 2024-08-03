package authen

import "github.com/themethaithian/go-pos-system/app"

type Handler interface {
	Login(ctx app.Context)
}

type handler struct {
	storage Storage
}

func NewHandler(storage Storage) Handler {
	return &handler{
		storage: storage,
	}
}
