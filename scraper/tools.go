package scraper

import (
	"fmt"
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
func getAvailabilityURL(o, d string, dd time.Time) string {
	return fmt.Sprintf(`https://www.logitravel.com/flightssaleprocess/availability?`+
		`spc=&ciaAerea=&ciasCon=&vueloslowcost=on&clase=0&descuentos=0`+
		`&hidOrigenSV=%v&hidDestinoSV=%v&searchTypeSV=%v`+
		`&fechaIdaSV=%v&fechaVueltaSV=%v&adultos=%v&ninos=%v&bebes=%v`,
		o, d, "OneWay", dd.Format("02/01/2006"), "", 1, 0, 0)
}

// diffDays get diff days from two dates
func diffDays(d1, d2 time.Time) int {
	return int(d2.Sub(d1) / (24 * time.Hour))
}
