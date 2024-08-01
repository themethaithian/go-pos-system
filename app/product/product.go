package product

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	ID          string
	Name        string
	Description *string
	Price       decimal.Decimal
	Quantity    int
}
