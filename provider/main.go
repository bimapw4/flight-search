package main

import (
	"context"
	"flight-api-provider/bootstrap"
	"flight-api-provider/internal/business"
	"flight-api-provider/internal/consumer"
	"flight-api-provider/internal/handlers"
	"flight-api-provider/internal/routes"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

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

	app.Use(requestid.New())

	// providercfg := bootstrap.Provider()
	// provider := provider.NewProvider(providercfg)

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	rdb := bootstrap.InitRedis(ctx)
	defer func() {
		log.Println("Redis Close")
		_ = rdb.Close()
	}()

	// repo := repositories.NewRepository(db)
	business := business.NewBusiness(rdb)
	handler := handlers.NewHandler(business)
	// middleware := middleware.NewAuthentication(business)

	consumer := consumer.NewConsumer(rdb, &business)
	go consumer.Run(ctx)

	routes.Routes(app, handler)

	go func() {
		port := ":3001"
		if os.Getenv("PORT") != "" {
			port = fmt.Sprintf(":%v", os.Getenv("PORT"))
		}
		log.Println(app.Listen(port))
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	if err := app.Shutdown(); err != nil {
		log.Printf("Fiber shutdown error: %v", err)
	} else {
		log.Println("Fiber shutdown complete")
	}
}
