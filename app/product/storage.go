package product

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type Storage interface {
	InsertProduct(id string, product NewProductRequest) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s *storage) InsertProduct(id string, product NewProductRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := `INSERT INTO tbl_product(id, name, description) VALUES($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, sql, id, product.Name, product.Description)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("cannot insert product: %s", product.Name))
	}

	return nil
}
