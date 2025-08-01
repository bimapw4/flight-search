package bootstrap

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func InitMeterProvider(ctx context.Context, serviceName string, app *fiber.App) (func(context.Context) error, error) {
	reader := metric.NewManualReader()

	provider := metric.NewMeterProvider(
		metric.WithReader(reader),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)

	otel.SetMeterProvider(provider)

	// ✅ Adaptor from net/http → fasthttp
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	log.Println("📊 Prometheus metrics available at /metrics")

	return provider.Shutdown, nil
}
