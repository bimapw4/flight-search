package routes

import (
	"flight-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handler handlers.Handlers) {
	// register route
	routes := []func(app *fiber.App, handler handlers.Handlers){
		NewFlight,
	}

	for _, route := range routes {
		route(app, handler)
	}
}
