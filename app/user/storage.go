package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Storage interface {
	InsertNewUser(user User) error
	UpdateUserRole(user AssignRole) error
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
		return fmt.Errorf("failed to insert user cause %v", err)
	}

	return nil
}

func (s *storage) UpdateUserRole(user AssignRole) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := `UPDATE tbl_user SET role = $1 WHERE id = $2`
	_, err := s.db.ExecContext(ctx, sql, user.Id, user.Role)
	if err != nil {
		return fmt.Errorf("failed to assign role for user (id: %d) cause %v", user.Id, err)
	}

	return nil
}
