package routes

import (
	"flight-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func NewFlight(app *fiber.App, handler handlers.Handlers) {
	app.Post("/api/flights/search", handler.FLight.PostSearchFlight)
	app.Get("/api/flights/search/:search_id/stream", handler.FLight.SSEFlightStream)
}
