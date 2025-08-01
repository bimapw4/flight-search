package business

import "github.com/redis/go-redis/v9"

type Business struct {
}

func NewBusiness(rdb *redis.Client) Business {
	return Business{}
}
