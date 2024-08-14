package hashing

type Hashing interface{}

type hashing struct{}

func NewHashing() Hashing {
	return &hashing{}
}
