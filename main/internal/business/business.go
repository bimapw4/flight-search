package business

import (
	"flight-api/internal/business/flight"

	"github.com/redis/go-redis/v9"
)

type Business struct {
	Flight flight.Business
}

func NewBusiness(rdb *redis.Client) Business {
	return Business{
		Flight: flight.NewBusiness(rdb),
	}
}
