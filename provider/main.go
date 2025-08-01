package main

import (
	"flight-api-provider/bootstrap"
	"flight-api-provider/internal/business"
	"flight-api-provider/internal/handlers"
	"flight-api-provider/internal/repositories"
	"flight-api-provider/internal/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	// Default config
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  os.Getenv("APP_NAME"),
		AppName:       os.Getenv("APP_NAME"),
	})

	app.Use(logger.New())

	// Connect to the PostgreSQL database
	db := bootstrap.ConnectDB()
	bootstrap.RunMigrations(db)

	app.Use(requestid.New())

	// providercfg := bootstrap.Provider()
	// provider := provider.NewProvider(providercfg)

	repo := repositories.NewRepository(db)
	business := business.NewBusiness(&repo)
	handler := handlers.NewHandler(business)
	// middleware := middleware.NewAuthentication(business)

	routes.Routes(app, handler)

	port := ":3000"
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%v", os.Getenv("PORT"))
	}

	log.Println(app.Listen(port))
}
