package db

import (
	"database/sql"
	"fmt"
	"price-tracker/logger"
	"time"

	// Postgres driver
	_ "github.com/lib/pq"
)

// db connection to database
var db *sql.DB

// DB Settings
const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "cguzman"
	dbPass = "cguzman"
	dbName = "price_tracker"
)

// InitDB Instantiate DB
func InitDB() {
	// Data connection params
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	var err error
	db, err = sql.Open("postgres", dataSourceName)

	if err != nil {
		logger.Panic(err)
	}
	if err = db.Ping(); err != nil {
		logger.Panic(err)
	}
}

// InsertPrice add row to Price db table
func InsertPrice(price float64, cia, flight, o, d string, d1 time.Time) {
	sqlQuery := `INSERT INTO public.price(amount, cia, flight_number, origin, destination, departure, comeback, oneway, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, null, true, now())`

	_, err := db.Exec(sqlQuery, fmt.Sprintf("%.2f", price), cia, flight, o, d, d1.Format("02/01/2006"))
	if err != nil {
		panic(err)
	}
}
