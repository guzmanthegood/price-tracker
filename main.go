package main

import (
	"price-tracker/logger"
	"price-tracker/scraper"
	"time"
)

func init() {
}

func main() {
	logger.Info("====== price-Tracker application Start ======")

	// Params Hardcode | TODO: take params from terminal
	origin := "PMI"
	destination := "EDI"

	d1, _ := time.Parse(scraper.DefaultDateFormat, "01/07/2018")
	d2, _ := time.Parse(scraper.DefaultDateFormat, "10/07/2018")

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
