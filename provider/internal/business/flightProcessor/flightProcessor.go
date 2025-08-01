package flightprocessor

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type FlightProcessor interface {
	Process(ctx context.Context, msg redis.XMessage) error
}

type business struct {
	rdb *redis.Client
}

func NewFlightProcessor(rdb *redis.Client) FlightProcessor {
	return &business{
		rdb: rdb,
	}
}

func (b *business) Process(ctx context.Context, msg redis.XMessage) error {
	values := msg.Values

	searchID, _ := values["search_id"].(string)
	origin, _ := values["from"].(string)    // fieldnya "from", bukan "origin"
	destination, _ := values["to"].(string) // sama
	date, _ := values["date"].(string)

	log.Printf("✈️ Processing: %s → %s (search_id=%s)", origin, destination, searchID)

	time.Sleep(2 * time.Second) // Simulasi delay API

	flight := map[string]interface{}{
		"id":             "flight-uuid-2",
		"airline":        "Lion Air",
		"flight_number":  "JT123",
		"from":           origin,
		"to":             destination,
		"departure_time": date + " 14:00",
		"arrival_time":   date + " 17:00",
		"price":          800000,
		"currency":       "IDR",
		"available":      true,
	}

	by, _ := json.Marshal([]map[string]interface{}{flight})

	// Hasilnya berupa array of flight(s)
	if _, err := b.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "flight.search.results",
		Values: map[string]interface{}{
			"search_id": searchID,
			"status":    "completed",
			"results":   by,
		},
	}).Result(); err != nil {
		log.Println("Failed to respond:", err)
		return err
	}

	// Acknowledge message
	if _, err := b.rdb.XAck(ctx, "flight.search.requested", "provider-group", msg.ID).Result(); err != nil {
		log.Printf(" Failed to XACK %s: %v", msg.ID, err)
		return err
	}

	log.Printf("Responded to %s", searchID)
	return nil
}
