package tracing

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Verify interface compliance for SpanExporter at compile time.
var _ sdktrace.SpanExporter = noopExporter{}

type noopExporter struct{}

// ExportSpans exports a batch of spans that perform no operations.
func (noopExporter) ExportSpans(context.Context, []sdktrace.ReadOnlySpan) error {
	return nil
}

// Shutdown notifies the exporter of a pending halt to operations.
func (noopExporter) Shutdown(context.Context) error {
	return nil
}

// newNoopTracerProvider configures a tracer provider that performs no operations with given resource.
func newNoopTracerProvider(res *resource.Resource) *sdktrace.TracerProvider {
	exporter := new(noopExporter)

	return sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(
			sdktrace.NewBatchSpanProcessor(exporter),
		),
	)
}
