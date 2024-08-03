package app

const (
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"
)

type Response struct {
	Status  string
	Message string
	Data    any
}
