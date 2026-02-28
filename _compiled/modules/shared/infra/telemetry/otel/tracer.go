package otel

import (
	"context"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	otelTrace "go.opentelemetry.io/otel/trace"
)

func setupTracer(ctx context.Context) (*trace.TracerProvider, *otelTrace.Tracer, error) {
	grpcExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(os.Getenv("OTEL_TRACER_ENDPOINT")),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, nil, err
	}
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(os.Getenv("SERVICE_NAME")),
		semconv.ServiceVersionKey.String(os.Getenv("WORKITO_VERSION")),
		attribute.String("env", os.Getenv("ENV")),
	)
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(grpcExporter,
			trace.WithBatchTimeout(time.Second), // Optional: configure batch timeout
		),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)
	tracer := otel.Tracer("workito")
	return tracerProvider, &tracer, nil
}
