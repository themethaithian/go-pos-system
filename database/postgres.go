package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/themethaithian/go-pos-system/config"
)

func NewPostgres() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=go-pos sslmode=disable", config.Val.DBHost, config.Val.DBPort, config.Val.DBUser, config.Val.DBPassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	return db
}
