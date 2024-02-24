package application

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432 // Default port for PostgreSQL
	user     = "postgres"
	password = "admin"
	dbname   = "evasion"
)

func Queryable() (*sql.DB, error) {
	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return sql.Open("postgres", connStr)
}
