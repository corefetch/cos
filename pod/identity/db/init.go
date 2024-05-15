package db

import (
	"database/sql"
	"gom/core/sys"
	"os"

	_ "github.com/lib/pq"
) // add this

var db *sql.DB

func Init() {

	var err error

	dburi := os.Getenv("DB")

	sys.Logger().Info("Database connecting...")

	db, err = sql.Open("postgres", dburi)

	if err != nil {
		sys.Logger().Fatalf("Failed to connect to %s", dburi)
	}
}
