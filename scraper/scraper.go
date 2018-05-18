package scraper

import (
	"fmt"
	"net"
	"net/http"
	"price-tracker/logger"
	"time"

	"github.com/gocolly/colly"
)

// DefaultDateFormat default format for spanish dates
const DefaultDateFormat = "02/01/2006"

// LoadAvailabilityPrices load prices in DB from flights availability page
func LoadAvailabilityPrices(o, d string, d1, d2 time.Time) error {
	logger.Info(fmt.Sprintf("Request -> Origin[%v] Destination[%v] DateRange[%v]..[%v]", o, d, d1.Format(DefaultDateFormat), d2.Format(DefaultDateFormat)))

	// New colly collector every iteration
	c := newCollyCollector()

	// Prices count
	var count int

	// Callbacks span price on availability
	c.OnHTML(".selectFlightOptionTrigger", func(e *colly.HTMLElement) {
		pr, _ := e.DOM.Attr("data-provider")
		fn, _ := e.DOM.Attr("data-number")
		va := e.DOM.Find("div.flight-price span.precioMedio span.price").Text()

		logger.Debug(fmt.Sprintf("Price Read -> Origin[%v] Destination[%v] Date[%v] provider[%v] flight[%v] amount[%v]",
			o, d, d1.Format(DefaultDateFormat), pr, fn, va))

		count++
	})

	c.OnScraped(func(r *colly.Response) {
		logger.Info(fmt.Sprintf("Result -> Origin[%v] Destination[%v] Date[%v] PricesFound[%v]", o, d, d1.Format(DefaultDateFormat), count))
	})

	// Next lvl of recursion
	if diffDays(d1, d2) > 0 {
		url := getAvailabilityURL(o, d, d1)
		logger.Debug("Visit -> ", url)
		c.Visit(url)
		LoadAvailabilityPrices(o, d, d1.AddDate(0, 0, 1), d2)
	}

	return nil
}

// newCollyCollector http config
func newCollyCollector() *colly.Collector {
	c := colly.NewCollector()

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       120 * time.Second,
		TLSHandshakeTimeout:   20 * time.Second,
		ExpectContinueTimeout: 2 * time.Second,
	})

	return c
}
