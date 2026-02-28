package telemetry

import (
	"context"
	"net/http"
	"_compiled/modules/shared/infra/telemetry/otel"
)

var telemetry Telemetry

type Telemetry interface {
	Shutdown(ctx context.Context) error
	LoggerMiddleware() func(next http.Handler) http.Handler
	TracerMiddleware() func(next http.Handler) http.Handler
	LogDebug(msg string, attrs any)
	LogInfo(msg string, attrs any)
	LogWarning(msg string, attrs any)
	LogError(msg string, err error, attrs any)
	WithSpan(
		ctx context.Context,
		spanName string,
		execute func(ctx context.Context) (any, error),
		attrs map[string]any,
	) (any, error)
}

func SetupTelemetry(ctx context.Context) (Telemetry, error) {
	if telemetry != nil {
		return telemetry, nil
	}
	tel, err := otel.NewOtelTelemetry(ctx)
	if err != nil {
		return nil, err
	}
	telemetry = tel
	return telemetry, nil
}

type attr struct {
	key   string
	value any
}

func Attr(key string, value any) attr {
	return attr{key, value}
}

func buildAttrMap(attrs []attr) map[string]any {
	attrsMap := make(map[string]any)
	for _, attr := range attrs {
		attrsMap[attr.key] = attr.value
	}
	return attrsMap
}

func LogDebug(msg string, attrs ...attr) {
	attrsMap := buildAttrMap(attrs)
	telemetry.LogDebug(msg, attrsMap)
}

func LogInfo(msg string, attrs ...attr) {
	attrsMap := buildAttrMap(attrs)
	telemetry.LogInfo(msg, attrsMap)
}

func LogWarning(msg string, attrs ...attr) {
	attrsMap := buildAttrMap(attrs)
	telemetry.LogWarning(msg, attrsMap)
}

func LogError(msg string, err error, attrs ...attr) {
	attrsMap := buildAttrMap(attrs)
	telemetry.LogError(msg, err, attrsMap)
}

func WithSpan[T any](
	ctx context.Context,
	spanName string,
	execute func(ctx context.Context) (T, error),
	attrs ...attr,
) (T, error) {
	attrMap := buildAttrMap(attrs)
	res, err := telemetry.WithSpan(ctx, spanName, func(spanCtx context.Context) (any, error) {
		return execute(spanCtx)
	}, attrMap)
	validRes, _ := res.(T)
	return validRes, err
}
