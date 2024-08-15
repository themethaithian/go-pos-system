package product

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	CategoryId  int    `json:"categoryId"`
	StockLevel  int    `json:"stockLevel"`
}
