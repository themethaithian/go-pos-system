package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Storage interface {
	InsertNewUser(user User) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) InsertNewUser(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := `INSERT INTO tbl_user(username, password_hash, role, created_at, updated_at) VALUES($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, sql, user.Username, user.PasswordHash, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("cannot insert user: %v", err)
	}

	return nil
}
