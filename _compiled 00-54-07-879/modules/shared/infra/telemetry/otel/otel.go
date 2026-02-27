package otel

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/trace"
	otelTrace "go.opentelemetry.io/otel/trace"
)

type OtelTelemetry struct {
	logger         *zerolog.Logger
	loggerProvider *log.LoggerProvider
	tracer         *otelTrace.Tracer
	tracerProvider *trace.TracerProvider
}

func NewOtelTelemetry(ctx context.Context) (*OtelTelemetry, error) {
	loggerProvider, logger, err := setupLogger(ctx)
	if err != nil {
		return nil, err
	}
	tracerProvider, tracer, err := setupTracer(ctx)
	if err != nil {
		return nil, err
	}
	return &OtelTelemetry{
		logger:         logger,
		loggerProvider: loggerProvider,
		tracerProvider: tracerProvider,
		tracer:         tracer,
	}, nil
}

func (o *OtelTelemetry) Shutdown(ctx context.Context) error {
	if err := o.loggerProvider.Shutdown(ctx); err != nil {
		return err
	}
	if err := o.tracerProvider.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func (o *OtelTelemetry) LoggerMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				t2 := time.Now()

				// Recover and record stack traces in case of a panic
				if rec := recover(); rec != nil {
					o.logger.Error().
						Str("type", "error").
						Timestamp().
						Interface("recover_info", rec).
						Bytes("debug_stack", debug.Stack()).
						Msg("log system error")
					http.Error(
						ww,
						http.StatusText(http.StatusInternalServerError),
						http.StatusInternalServerError,
					)
				}

				// log end request
				o.logger.Info().
					Str("type", "access").
					Timestamp().
					Fields(map[string]interface{}{
						"remote_ip":  r.RemoteAddr,
						"url":        r.URL.Path,
						"proto":      r.Proto,
						"method":     r.Method,
						"user_agent": r.Header.Get("User-Agent"),
						"status":     ww.Status(),
						"latency_ms": float64(t2.Sub(t1).Nanoseconds()) / 1000000.0,
						"bytes_in":   r.Header.Get("Content-Length"),
						"bytes_out":  ww.BytesWritten(),
					}).
					Msg(fmt.Sprintf("incoming_request %s", r.URL.Path))
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}

func (o *OtelTelemetry) TracerMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Do not add tracing for static content
			if strings.HasPrefix(r.URL.Path, "/public") {
				next.ServeHTTP(w, r)
				return
			}
			ctx := r.Context()
			spanCtx, span := (*o.tracer).Start(ctx, r.URL.Path)
			defer span.End()
			_r := r.WithContext(spanCtx)
			next.ServeHTTP(w, _r)
		}
		return http.HandlerFunc(fn)
	}
}

func (o *OtelTelemetry) LogDebug(msg string, attrs any) {
	o.logger.Debug().Fields(attrs).Msg(msg)
}

func (o *OtelTelemetry) LogInfo(msg string, attrs any) {
	o.logger.Info().Fields(attrs).Msg(msg)
}

func (o *OtelTelemetry) LogWarning(msg string, attrs any) {
	o.logger.Warn().Fields(attrs).Msg(msg)
}

func (o *OtelTelemetry) LogError(msg string, err error, attrs any) {
	o.logger.Error().Err(err).Msg(msg)
}

func (o *OtelTelemetry) WithSpan(
	ctx context.Context,
	spanName string,
	execute func(spanCtx context.Context) (any, error),
	attrs map[string]any,
) (any, error) {
	attributes := make([]attribute.KeyValue, 0)
	for key, value := range attrs {
		switch v := value.(type) {
		case string:
			attributes = append(attributes, attribute.String(key, v))
		case bool:
			attributes = append(attributes, attribute.Bool(key, v))
		case int:
			attributes = append(attributes, attribute.Int(key, v))
		case float64:
			attributes = append(attributes, attribute.Float64(key, v))
		default:
			attributes = append(attributes, attribute.String(key, fmt.Sprintf("%s", v)))
		}
	}
	spanCtx, span := (*o.tracer).Start(ctx, spanName, otelTrace.WithAttributes(attributes...))
	defer span.End()
	res, err := execute(spanCtx)
	if err != nil {
		span.RecordError(err)
	}
	return res, err
}
