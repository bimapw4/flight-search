// bootstrap/redis.go
package bootstrap

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func InitRedis(ctx context.Context) *redis.Client {
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       redisDB,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	log.Println("Redis Connected Successfully!")

	return rdb
}
