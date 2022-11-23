package tracing

import (
	"context"
	"fmt"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// Initialize initializes an OpenTelemetry-compatible tracing implementation.
func Initialize(ctx context.Context, config *Config) (func(ctx context.Context) error, error) {
	if config == nil {
		return nil, fmt.Errorf("tracing: empty config provided")
	}

	if !config.Enabled {
		return newNoopTracerProvider(resource.Default()).Shutdown, nil
	}

	// Create a new resource with details of the service.
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(config.ServiceName),
		semconv.ServiceVersionKey.String(config.ServiceVersion),
		semconv.DeploymentEnvironmentKey.String(config.Environment),
	)

	var tp *sdktrace.TracerProvider
	var err error

	switch config.Provider {
	case "elastic":
		tp, err = newElasticTracerProvider(ctx, res, config.Elastic)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize Elastic tracer provider: %w", err)
		}
	case "otel_collector":
		tp, err = newOTelCollectorTracerProvider(ctx, res, config.OTelCollector)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize OTel Collector tracer provider: %w", err)
		}
	case "jaeger":
		tp, err = newJaegerTracerProvider(ctx, res, config.Jaeger)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize Jaeger tracer provider: %w", err)
		}
	case "noop":
		tp = newNoopTracerProvider(res)
	default:
		return nil, fmt.Errorf("tracing: invalid provider %s", config.Provider)
	}

	// Register tp as the global tracer provider.
	otel.SetTracerProvider(tp)

	// Set global propagator to composition of baggage and tracecontext.
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.Baggage{},
			propagation.TraceContext{},
		),
	)

	// Shutdown will flush any remaining spans and shut down the exporter.
	return tp.Shutdown, nil
}

// removeURLSchemePrefix prepares endpoint for OTel devoid of URL scheme.
func removeURLSchemePrefix(url string) string {
	return strings.ReplaceAll(url, "https://", "")
}
