package authen

import "database/sql"

type Storage interface {
	CreateUser(register Register) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreateUser(register Register) error {
	return nil
}
