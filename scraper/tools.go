package scraper

import (
	"net/url"
	"time"
)

// Price includes basic flight data
type Price struct {
	Amount      float64
	FlighNumber string
	Origin      string
	Destination string
	Departure   string
	Comeback    string
	OneWay      string
}

// getAvailabilityURL Logitravel availability creation
func getAvailabilityURL(o, d string, d1 time.Time) string {
	str, _ := url.Parse("https://www.logitravel.com")
	str.Path += "/flightssaleprocess/availability"
	parameters := url.Values{}
	parameters.Add("vueloslowcost", "on")
	parameters.Add("clase", "0")
	parameters.Add("descuentos", "0")
	parameters.Add("hidOrigenSV", o)
	parameters.Add("hidDestinoSV", d)
	parameters.Add("searchTypeSV", "OneWay")
	parameters.Add("fechaIdaSV", d1.Format("02/01/2006"))
	parameters.Add("fechaVueltaSV", "")
	parameters.Add("adultos", "1")
	parameters.Add("ninos", "")
	parameters.Add("bebes", "")
	str.RawQuery = parameters.Encode()
	return str.String()
}

// getAvailabilityURL Logitravel availability creation
func getFilteredAvailabilityURL(o, d string, d1 time.Time, idCache string) string {
	var hashItineraries = o + "." + d + "." + d1.Format("2006-01-02")

	str, _ := url.Parse("https://www.logitravel.com")
	str.Path += "/flightssaleprocess/Filter/GetFilteredAvailability"
	parameters := url.Values{}
	parameters.Add("idCache", idCache)
	parameters.Add("stopOver", "0,1")
	parameters.Add("companies", "")
	parameters.Add("allCompaniesSelected", "true")
	parameters.Add("airports", "")
	parameters.Add("sortCriteria", "0")
	parameters.Add("discount", "0")
	parameters.Add("currentPage", "2")
	parameters.Add("hours", "0,23;0,23;")
	parameters.Add("hashItineraries", hashItineraries)
	parameters.Add("includeSharedCodes", "true")
	str.RawQuery = parameters.Encode()
	return str.String()
}

// diffDays get diff days from two dates
func diffDays(d1, d2 time.Time) int {
	return int(d2.Sub(d1) / (24 * time.Hour))
}
