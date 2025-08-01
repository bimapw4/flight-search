package consumer

import (
	"context"
	"flight-api-provider/internal/business"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Consumer struct {
	rdb      *redis.Client
	business *business.Business
}

func NewConsumer(rdb *redis.Client, business *business.Business) *Consumer {
	return &Consumer{
		rdb:      rdb,
		business: business,
	}
}

func (c *Consumer) Run(ctx context.Context) {
	err := c.rdb.XGroupCreateMkStream(ctx, "flight.search.requested", "provider-group", "$")
	if err != nil {
		log.Println(" Failed to create group: %v", err)
	}

	log.Println("Provider consumer listening on flight.search.requested...")

	for {
		streams, err := c.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "provider-group",
			Consumer: "provider-1",
			Streams:  []string{"flight.search.requested", ">"},
			Block:    5 * time.Second,
			Count:    1,
		}).Result()

		if err != nil && err != redis.Nil {
			log.Println("⚠️ Error reading group:", err)
			continue
		}

		for _, stream := range streams {
			for _, msg := range stream.Messages {
				go func(msg redis.XMessage) {
					err := c.business.Flightprocessor.Process(ctx, msg)
					if err == nil {
						_, ackErr := c.rdb.XAck(ctx, "flight.search.requested", "provider-group", msg.ID).Result()
						if ackErr != nil {
							log.Printf("⚠️ Failed to XACK: %v", ackErr)
						}
					}
				}(msg)
			}
		}
	}
}
