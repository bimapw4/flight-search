package routes

import (
	"flight-api-provider/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handler handlers.Handlers) {
	// register route
	routes := []func(app *fiber.App, handler handlers.Handlers){}

	for _, route := range routes {
		route(app, handler)
	}
}
