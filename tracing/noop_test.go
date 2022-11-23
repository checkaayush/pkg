package tracing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/resource"
)

func TestNewNoopTracerProvider(t *testing.T) {

	assert := require.New(t)

	res := resource.Default()
	tp := newNoopTracerProvider(res)
	assert.NotNil(tp)
}

func TestNoopExporterExportSpans(t *testing.T) {

	assert := require.New(t)

	exp := new(noopExporter)
	err := exp.ExportSpans(context.TODO(), nil)
	assert.Nil(err)
}

func TestNoopExporterShutdown(t *testing.T) {

	assert := require.New(t)

	exp := new(noopExporter)
	err := exp.Shutdown(context.TODO())
	assert.Nil(err)
}
