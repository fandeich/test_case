package postgres

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitPostgres(_ context.Context, connString string) *sql.DB {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
