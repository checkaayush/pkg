package tracing

import (
	"context"
	"fmt"
	"strings"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// newOTelCollectorTracerProvider configures a new tracer provider with given resource for OTel Collector.
// Refer https://opentelemetry.io/docs/collector for more details.
func newOTelCollectorTracerProvider(ctx context.Context, res *resource.Resource, config *OTelCollectorConfig) (*sdktrace.TracerProvider, error) {
	if config == nil {
		return nil, fmt.Errorf("otel collector config empty")
	}

	exporter, err := newOTelCollectorExporter(ctx, config)
	if err != nil {
		return nil, err
	}

	// Create a new and configured tracer provider.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(config.TransactionSampling)),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(
			sdktrace.NewBatchSpanProcessor(exporter),
		),
	)
	return tp, nil
}

// newOTelCollectorExporter returns a protocol-aware exporter for OTel Collector.
func newOTelCollectorExporter(ctx context.Context, config *OTelCollectorConfig) (*otlptrace.Exporter, error) {
	if config.ServerURL == "" {
		return nil, fmt.Errorf("invalid OTel collector URL")
	}

	endpoint := removeURLSchemePrefix(config.ServerURL)

	if strings.EqualFold(config.Protocol, "grpc") {
		opts := []otlptracegrpc.Option{
			otlptracegrpc.WithEndpoint(endpoint),
			otlptracegrpc.WithInsecure(),
		}
		if config.SecretToken != "" {
			opts = append(opts, otlptracegrpc.WithHeaders(map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", config.SecretToken),
			}))
		}
		return otlptracegrpc.New(ctx, opts...)
	}

	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(endpoint),
	}
	if config.SecretToken != "" {
		opts = append(opts, otlptracehttp.WithHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", config.SecretToken),
		}))
	}
	return otlptracehttp.New(ctx, opts...)
}
