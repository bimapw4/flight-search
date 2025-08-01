package flight

import (
	"bufio"
	"context"
	"encoding/json"
	"flight-api/internal/business"
	"flight-api/internal/entity"
	"flight-api/internal/response"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Handler interface {
	PostSearchFlight(c *fiber.Ctx) error
	SSEFlightStream(c *fiber.Ctx) error
}

type handler struct {
	business business.Business
	rdb      *redis.Client
}

func NewHandler(business business.Business, rdb *redis.Client) Handler {
	return &handler{
		business: business,
		rdb:      rdb,
	}
}
func (h *handler) PostSearchFlight(c *fiber.Ctx) error {

	var (
		Entity = "SearchFlight"
	)

	var input entity.FlightSearchInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	res, err := h.business.Flight.PublishSearch(c.UserContext(), input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to publish to stream"})
	}

	return response.NewResponse(Entity).
		Success("Search request submitted", res).
		JSON(c, fiber.StatusOK)
}

func (h *handler) SSEFlightStream(c *fiber.Ctx) error {
	searchID := c.Params("search_id")
	if searchID == "" {
		return c.Status(400).SendString("Missing search_id")
	}

	c.Set("Content-Type", "text/event-stream")
	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		lastID := "0"

		ctx, _ := context.WithTimeout(c.Context(), 5*time.Second)

		for {
			streams, err := h.rdb.XRead(ctx, &redis.XReadArgs{
				Streams: []string{"flight.search.results", lastID},
				Block:   5 * time.Second,
			}).Result()

			if err != nil && err != redis.Nil {
				continue
			}

			for _, stream := range streams {
				for _, msg := range stream.Messages {
					lastID = msg.ID
					if val, ok := msg.Values["search_id"]; ok && val == searchID {
						data, _ := json.Marshal(msg.Values)
						fmt.Fprintf(w, "data: %s\n\n", data)
						w.Flush()
					}
				}
			}
		}
	})

	return nil
}
