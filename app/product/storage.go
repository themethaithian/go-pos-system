package product

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Storage interface {
	RetrieveAllProducts() ([]Product, error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s *storage) RetrieveAllProducts() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	products := make([]Product, 0, 100)

	sql := `SELECT id, name, description, price, category_id, stock_level FROM tbl_product`

	rows, err := s.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products cause %v", err)
	}
	defer rows.Close()

	rowCount := 1

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.StockLevel)
		if err != nil {
			return nil, fmt.Errorf("failed to scan on row %d cause %v", rowCount, err)
		}

		products = append(products, product)

		rowCount++
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("an error occur when retrieving products cause %v", err)
	}

	return products, nil
}
