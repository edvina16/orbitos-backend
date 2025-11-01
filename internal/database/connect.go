package database

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Connect opens a database connection for use with sqlc's generated code.
func Connect() (*sql.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://orbitos:orbitospass@localhost:5433/orbitos"
	}
	return sql.Open("pgx", dbUrl)
}
