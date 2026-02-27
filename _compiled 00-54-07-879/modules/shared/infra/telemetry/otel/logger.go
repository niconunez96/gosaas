package otel

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/bridges/otelzerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func setupLogger(ctx context.Context) (*log.LoggerProvider, *zerolog.Logger, error) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	grpcExport, err := otlploggrpc.New(ctx,
		// Use WithInsecure for development/testing if collector has no TLS configured
		// For production, ensure secure connection
		otlploggrpc.WithInsecure(),
		otlploggrpc.WithEndpoint(os.Getenv("OTEL_LOGGER_ENDPOINT")),
	)
	if err != nil {
		return nil, nil, err
	}
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(os.Getenv("SERVICE_NAME")),
		semconv.ServiceVersionKey.String(os.Getenv("SERVICE_VERSION")),
		attribute.String("env", os.Getenv("ENV")),
	)
	processor := log.NewBatchProcessor(grpcExport)
	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(processor),
		log.WithResource(res),
	)
	hook := otelzerolog.NewHook("zerologhook", otelzerolog.WithLoggerProvider(loggerProvider))
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Hook(hook)
	return loggerProvider, &logger, nil
}
