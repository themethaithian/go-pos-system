package user

type Handler interface{}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
}
