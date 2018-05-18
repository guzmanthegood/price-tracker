package db

import (
	"database/sql"
	"price-tracker/logger"

	// Postgres driver
	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB Instantiate DB
func InitDB() {
	var err error
	db, err := sql.Open("postgres", "user=cguzman password=cguzman dbname=price_tracker sslmode=disable")
	if err != nil {
		logger.Panic(err)
	}

	if err = db.Ping(); err != nil {
		logger.Panic(err)
	}
}
