package authen

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Storage interface {
	RetrieveRoleFromUsername(username string) (string, error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) RetrieveRoleFromUsername(username string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var role string

	sql := `SELECT role FROM tbl_user WHERE username = $1`

	if err := s.db.QueryRowContext(ctx, sql, username).Scan(&role); err != nil {
		return "", fmt.Errorf("failed to retrive role from %s cause %v", username, err)
	}

	return role, nil
}
