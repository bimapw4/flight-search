package main

import (
	"context"
	"flight-api/bootstrap"
	"flight-api/internal/business"
	"flight-api/internal/handlers"
	"flight-api/internal/routes"
	"fmt"
	"log"
	"os"
	"time"

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
	// db := bootstrap.ConnectDB()
	// bootstrap.RunMigrations(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb := bootstrap.InitRedis(ctx)

	app.Use(requestid.New())

	// providercfg := bootstrap.Provider()
	// provider := provider.NewProvider(providercfg)

	// repo := repositories.NewRepository(db)
	business := business.NewBusiness(rdb)
	handler := handlers.NewHandler(business, rdb)
	// middleware := middleware.NewAuthentication(business)

	routes.Routes(app, handler)

	port := ":3000"
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%v", os.Getenv("PORT"))
	}

	log.Println(app.Listen(port))
}
