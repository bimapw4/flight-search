package business

import (
	flightprocessor "flight-api-provider/internal/business/flightProcessor"

	"github.com/redis/go-redis/v9"
)

type Business struct {
	Flightprocessor flightprocessor.FlightProcessor
}

func NewBusiness(rdb *redis.Client) Business {
	return Business{
		Flightprocessor: flightprocessor.NewFlightProcessor(rdb),
	}
}
