package flight

import (
	"context"
	"flight-api/internal/entity"
	"flight-api/internal/presentations"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Business interface {
	PublishSearch(ctx context.Context, req entity.FlightSearchInput) (*presentations.SearchFlight, error)
}

type business struct {
	rdb *redis.Client
}

func NewBusiness(rdb *redis.Client) Business {
	return &business{
		rdb: rdb,
	}
}
func (b *business) PublishSearch(ctx context.Context, req entity.FlightSearchInput) (*presentations.SearchFlight, error) {
	if req.SearchID == "" {
		req.SearchID = uuid.NewString()
	}

	_, err := b.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "flight.search.requested",
		Values: map[string]interface{}{
			"search_id":  req.SearchID,
			"from":       req.From,
			"to":         req.To,
			"date":       req.Date,
			"passengers": req.Passengers,
		},
	}).Result()

	return &presentations.SearchFlight{
		SearchID: req.SearchID,
		Status:   "processing",
	}, err
}
