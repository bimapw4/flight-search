package handlers

import (
	"flight-api/internal/business"
	"flight-api/internal/handlers/flight"

	"github.com/redis/go-redis/v9"
)

type Handlers struct {
	FLight flight.Handler
}

func NewHandler(business business.Business, rdb *redis.Client) Handlers {
	return Handlers{
		FLight: flight.NewHandler(business, rdb),
	}
}
