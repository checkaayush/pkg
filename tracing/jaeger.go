package tracing

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// newJaegerTracerProvider returns a new OTLP tracer provider for Jaeger.
func newJaegerTracerProvider(ctx context.Context, res *resource.Resource, config *JaegerConfig) (*sdktrace.TracerProvider, error) {
	if config == nil {
		return nil, fmt.Errorf("jaeger config empty")
	}

	exporter, err := newJaegerExporter(ctx, config)
	if err != nil {
		return nil, err
	}

	// Create a new and configured tracer provider.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(
			sdktrace.NewBatchSpanProcessor(exporter),
		),
	)
	return tp, nil
}

// newJaegerExporter returns an OTLP Jaeger exporter.
func newJaegerExporter(ctx context.Context, config *JaegerConfig) (*jaeger.Exporter, error) {
	if config == nil {
		return nil, fmt.Errorf("jaeger config empty")
	}

	endpoint := fmt.Sprintf("http://%s:%d/api/traces", config.Host, config.Port)
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
}
