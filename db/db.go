package db

import (
	"database/sql"
	"fmt"
	"price-tracker/logger"
	"time"

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

// InsertPrice add row to Price db table
func InsertPrice(price float64, flight, o, d string, d1 time.Time) {
	sqlQuery := `
	INSERT INTO public.price(amount, flight_number, origin, destination, departure, comeback, oneway, created_at)
	VALUES ($1, $2, $3, $4, $5, null, true, now())`

	_, err := db.Exec(sqlQuery, price, flight, o, d, d1.Format("02/01/2006"))
	if err != nil {
		panic(err)
	}

	//VALUES (120.0, 'FR1111', 'PMI', 'MAD', '01/01/2018', null, true);
}
