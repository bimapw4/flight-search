package entity

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type FlightSearchInput struct {
	SearchID   string `json:"search_id"`
	From       string `json:"from"`
	To         string `json:"to"`
	Date       string `json:"date"`
	Passengers int    `json:"passengers"`
}

func (v *FlightSearchInput) Validation() error {
	return validation.ValidateStruct(
		v,
		validation.Field(&v.SearchID, is.UUID),
		validation.Field(&v.From, validation.Required, validation.Length(2, 5), is.UpperCase),
		validation.Field(&v.To, validation.Required, validation.Length(2, 5), is.UpperCase),
		validation.Field(&v.Date, validation.Required, validation.Match(regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`)).Error("Must be valid format YYYY-MM-DD")),
		validation.Field(&v.Passengers, validation.Required, validation.Min(1), validation.Max(9)),
	)
}

type FlightResult struct {
	SearchID      string `json:"search_id"`
	ID            string `json:"id"`
	Airline       string `json:"airline"`
	FlightNumber  string `json:"flight_number"`
	From          string `json:"from"`
	To            string `json:"to"`
	DepartureTime string `json:"departure_time"`
	ArrivalTime   string `json:"arrival_time"`
	Price         int    `json:"price"`
	Currency      string `json:"currency"`
	Available     bool   `json:"available"`
}
