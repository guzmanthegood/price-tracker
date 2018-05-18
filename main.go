package main

import (
	"os"
	"price-tracker/logger"
	"price-tracker/scraper"
	"time"
)

func init() {
}

func main() {
	logger.Info("====== price-Tracker application Start ======")
	logger.Info("args: ", os.Args[1:])
	args := os.Args[1:]

	// Params Hardcode | TODO: take params from terminal
	origin := args[0]
	destination := args[1]

	d1, _ := time.Parse(scraper.DefaultDateFormat, args[2])
	d2, _ := time.Parse(scraper.DefaultDateFormat, args[3])

	scraper.LoadAvailabilityPrices(origin, destination, d1, d2)

	logger.Info("===== price-Tracker application End   ======")
}
