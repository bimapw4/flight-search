package flightprocessor

import (
	"context"
	"encoding/json"
	"flight-api-provider/internal/entity"
	"io/ioutil"
	"log"
	"strings"
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
	origin, _ := values["from"].(string)
	destination, _ := values["to"].(string)
	date, _ := values["date"].(string)
	passengers, _ := values["passengers"].(int)

	log.Printf("✈️ Processing: %s → %s (search_id=%s)", origin, destination, searchID)

	if _, err := b.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "flight.search.results",
		Values: map[string]interface{}{
			"search_id": searchID,
			"status":    "processing",
		},
	}).Result(); err != nil {
		log.Println("Failed to respond:", err)
		return err
	}

	time.Sleep(2 * time.Second)

	flights, err := b.searchJson(entity.FlightSearchInput{
		From:       origin,
		To:         destination,
		Date:       date,
		Passengers: passengers,
	})
	if err != nil {
		return err
	}

	if _, err := b.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "flight.search.results",
		Values: map[string]interface{}{
			"search_id": searchID,
			"status":    "completed",
			"results":   b.toJSON(flights),
		},
	}).Result(); err != nil {
		log.Println("Failed to respond:", err)
		return err
	}

	time.Sleep(1 * time.Second)

	if _, err := b.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "flight.search.results",
		Values: map[string]interface{}{
			"search_id":     searchID,
			"status":        "completed",
			"total_results": len(flights),
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

func (b *business) searchJson(search entity.FlightSearchInput) ([]entity.FlightResult, error) {

	body, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		log.Println("unable to read file: %v", err)
		return nil, err
	}

	simpleJson := []entity.FlightResult{}

	err = json.Unmarshal(body, &simpleJson)
	if err != nil {
		return nil, err
	}

	matched := []entity.FlightResult{}
	for _, v := range simpleJson {
		if v.Available &&
			strings.EqualFold(v.From, search.From) &&
			strings.EqualFold(v.To, search.To) &&
			strings.HasPrefix(v.DepartureTime, search.Date) {
			matched = append(matched, v)
		}
	}

	return matched, nil
}

func (b *business) toJSON(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
