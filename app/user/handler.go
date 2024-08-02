package user

type Handler interface{}

type handler struct {
	storage Storage
}

func NewHandler(storage Storage) Handler {
	return &handler{
		storage: storage,
	}
}
