package product

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Storage interface {
	InsertProduct(id string, newProduct NewProduct) error
	UpdateProduct(id string, editProduct EditProduct) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s *storage) InsertProduct(id string, newProduct NewProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := `INSERT INTO tbl_product(id, name, description) VALUES($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, sql, id, newProduct.Name, newProduct.Description)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("cannot insert product: %s", newProduct.Name))
	}

	return nil
}

func (s *storage) UpdateProduct(id string, editProduct EditProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := `UPDATE tbl_product SET `
	var args []interface{}
	argCount := 1

	if editProduct.Name != nil {
		sql += fmt.Sprintf("name = $%d, ", argCount)
		args = append(args, *editProduct.Name)
		argCount++
	}

	if editProduct.Description != nil {
		sql += fmt.Sprintf("description = $%d, ", argCount)
		args = append(args, *editProduct.Description)
		argCount++
	}

	if editProduct.Price != nil {
		sql += fmt.Sprintf("price = $%d, ", argCount)
		args = append(args, *editProduct.Price)
		argCount++
	}

	if editProduct.Quantity != nil {
		sql += fmt.Sprintf("quantity = $%d, ", argCount)
		args = append(args, *editProduct.Quantity)
		argCount++
	}

	sql = strings.TrimSuffix(sql, ", ")
	sql += fmt.Sprintf(" WHERE id = $%d", argCount)
	args = append(args, id)

	_, err := s.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("cannot edit product: %s", id))
	}

	return nil
}
