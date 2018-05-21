package db

import (
	"database/sql"
	"fmt"
	"price-tracker/logger"

	// Postgres driver
	_ "github.com/lib/pq"
)

// DB Settings
const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "cguzman"
	dbPass = "cguzman"
	dbName = "price_tracker"
)

var db *sql.DB

// InitDB Instantiate DB
func InitDB() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		logger.Panic(err)
	}

	if err = db.Ping(); err != nil {
		logger.Panic(err)
	}

	defer db.Close()
}
