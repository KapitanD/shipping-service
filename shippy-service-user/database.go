package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewConnection returns a new database connection instance
func NewConnection() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	conn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, host, dbName)
	log.Println(conn)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
