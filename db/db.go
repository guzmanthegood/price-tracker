package db

import (
	"database/sql"
	"fmt"
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
	}
	if err = db.Ping(); err != nil {
	}
}

// DeleteOldPrices delete old prices to replace (dangerous!!)
func DeleteOldPrices(o, d string, d1, d2 time.Time) {
	sqlQuery := `DELETE FROM price where origin = $1 AND destination = $2 AND departure >= $3 AND departure <=  $4`

	_, err := db.Exec(sqlQuery, o, d, d1.Format("02/01/2006"), d2.Format("02/01/2006"))
	if err != nil {
		panic(err)
	}
}

// InsertPrice add row to Price db table
func InsertPrice(price float64, cia, flight, o, d string, d1 time.Time) {
	sqlQuery := `INSERT INTO price(amount, cia, flight_number, origin, destination, departure, comeback, oneway, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, null, true, now())`

	_, err := db.Exec(sqlQuery, price, cia, flight, o, d, d1.Format("02/01/2006"))
	if err != nil {
		panic(err)
	}
}
