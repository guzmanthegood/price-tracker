package scraper

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func parseFlightPrice(e *colly.HTMLElement) Price {
	p, _ := parseEuroPrice(e.Text)
	return Price{
		Amount: p,
	}
}

//parseEuroPrice parse string price to float64
func parseEuroPrice(str string) (float64, error) {
	// Remove whitespaces, euro char, replace comma
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "€", "", -1)
	str = strings.Replace(str, ",", ".", -1)
	str = strings.Replace(str, "\u00a0", "", -1)
	str = strings.Replace(str, "\u20AC", "", -1)

	return strconv.ParseFloat(str, 64)
}
