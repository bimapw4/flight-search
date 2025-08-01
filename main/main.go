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

	shutdownMetrics, err := bootstrap.InitMeterProvider(ctx, "main-service", app)
	if err != nil {
		log.Fatalf("failed to init OTel metrics: %v", err)
	}
	defer func() {
		if err := shutdownMetrics(ctx); err != nil {
			log.Printf("failed to shutdown meter provider: %v", err)
		}
	}()

	rdb := bootstrap.InitRedis(ctx)
	defer func() {
		log.Println("Redis Close")
		_ = rdb.Close()
	}()

	app.Use(requestid.New())

	// providercfg := bootstrap.Provider()
	// provider := provider.NewProvider(providercfg)

	// repo := repositories.NewRepository(db)
	business := business.NewBusiness(rdb)
	handler := handlers.NewHandler(business, rdb)
	// middleware := middleware.NewAuthentication(business)

	routes.Routes(app, handler)

	go func() {
		port := ":3000"
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
