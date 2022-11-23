package tracing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/resource"
)

func TestNewJaegerExporterSuccess(t *testing.T) {

	assert := require.New(t)

	cfg := &JaegerConfig{
		Host: "localhost",
		Port: 14268,
	}
	exporter, err := newJaegerExporter(context.TODO(), cfg)
	assert.NotNil(exporter)
	assert.Nil(err)
}

func TestNewJaegerExporterFailure(t *testing.T) {

	assert := require.New(t)

	exporter, err := newJaegerExporter(context.TODO(), nil)
	assert.Nil(exporter)
	assert.NotNil(err)
}

func TestNewJaegerTracerProviderSuccess(t *testing.T) {

	assert := require.New(t)

	res := resource.Default()
	cfg := &JaegerConfig{
		Host: "localhost",
		Port: 14268,
	}
	tp, err := newJaegerTracerProvider(context.TODO(), res, cfg)
	assert.Nil(err)
	assert.NotNil(tp)
}

func TestNewJaegerTracerProviderFailure(t *testing.T) {

	assert := require.New(t)

	res := resource.Default()
	tp, err := newJaegerTracerProvider(context.TODO(), res, nil)
	assert.NotNil(err)
	assert.Nil(tp)
}
