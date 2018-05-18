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

/*

input := "2017-08-31"
layout := "2006-01-02"
t, _ := time.Parse(layout, input)
fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
fmt.Println(t.Format("02-Jan-2006")) // 31-Aug-2017

*/
