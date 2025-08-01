package entity

type FlightSearchInput struct {
	SearchID   string `json:"search_id"`
	From       string `json:"from"`
	To         string `json:"to"`
	Date       string `json:"date"`
	Passengers int    `json:"passengers"`
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
