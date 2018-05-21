package scraper

import (
	"fmt"
	"net"
	"net/http"
	"price-tracker/db"
	"price-tracker/logger"
	"time"

	"github.com/gocolly/colly"
)

// DefaultDateFormat default format for spanish dates
const DefaultDateFormat = "02/01/2006"

// LoadAvailabilityPrices load prices in DB from flights availability page
func LoadAvailabilityPrices(o, d string, d1, d2 time.Time) error {
	logger.Info(fmt.Sprintf("LoadAvailabilityPrices -> Origin[%v] Destination[%v] DateRange[%v]..[%v]",
		o, d, d1.Format(DefaultDateFormat), d2.Format(DefaultDateFormat)))

	// Init database
	db.InitDB()

	// Clear old results
	db.DeleteOldPrices(o, d, d1, d2)

	// New colly collector every iteration
	c := newCollyCollector()

	// Prices count
	var count int

	// Callbacks span price on availability
	c.OnHTML(".selectFlightOptionTrigger", func(e *colly.HTMLElement) {
		pr, _ := e.DOM.Attr("data-provider")
		fn, _ := e.DOM.Attr("data-number")
		va, _ := parseEuroPrice(e.DOM.Find("div.flight-price span.precioMedio span.price").Text())

		logger.Debug(fmt.Sprintf("Price Read -> Origin[%v] Destination[%v] Date[%v] provider[%v] flight[%v] amount[%v]",
			o, d, d1.Format(DefaultDateFormat), pr, fn, va))

		// Insert price in database
		db.InsertPrice(va, pr, fn, o, d, d1, getAvailabilityURL(o, d, d1))

		count++
	})

	c.OnScraped(func(r *colly.Response) {
		logger.Info(fmt.Sprintf("Result -> Origin[%v] Destination[%v] Date[%v] PricesFound[%v]", o, d, d1.Format(DefaultDateFormat), count))
	})

	// Next lvl of recursion
	if diffDays(d1, d2) >= 0 {
		idCache := getAvailabilityCacheID(o, d, d1)
		url := getFilteredAvailabilityURL(o, d, d1, idCache)
		c.Visit(url)

		LoadAvailabilityPrices(o, d, d1.AddDate(0, 0, 1), d2)
	}

	return nil
}

// getAvailabilityCacheID get cache id from availability page
func getAvailabilityCacheID(o, d string, d1 time.Time) string {
	logger.Debug(fmt.Sprintf("getAvailabilityCacheID -> Origin[%v] Destination[%v] Date[%v]", o, d, d1.Format(DefaultDateFormat)))

	c := newCollyCollector()

	var idCache string
	c.OnHTML("form#continue input[name=idCache]", func(e *colly.HTMLElement) {
		idCache, _ = e.DOM.Attr("value")
	})

	url := getAvailabilityURL(o, d, d1)
	c.Visit(url)

	return idCache
}

// newCollyCollector http config
func newCollyCollector() *colly.Collector {
	c := colly.NewCollector()

	c.WithTransport(

		&http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 60 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       100 * time.Second,
			TLSHandshakeTimeout:   20 * time.Second,
			ExpectContinueTimeout: 2 * time.Second,
		})

	return c
}
