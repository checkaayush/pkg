package tracing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/sdk/resource"
)

func TestNewOTelCollectorExporterSuccess(t *testing.T) {

	assert := require.New(t)

	cfg := &OTelCollectorConfig{
		ServerURL:           "otelcollector.plivops.com:443",
		SecretToken:         "token",
		TransactionSampling: 0.5,
		Protocol:            "http",
	}
	exporter, err := newOTelCollectorExporter(context.TODO(), cfg)
	assert.NotNil(exporter)
	assert.Nil(err)
}

func TestNewOTelCollectorExporterFailure(t *testing.T) {

	assert := require.New(t)

	cfg := &OTelCollectorConfig{
		ServerURL:           "",
		SecretToken:         "token",
		TransactionSampling: 0.5,
		Protocol:            "http",
	}
	exporter, err := newOTelCollectorExporter(context.TODO(), cfg)
	assert.Nil(exporter)
	assert.NotNil(err)
}

func TestNewOTelCollectorTracerProviderSuccess(t *testing.T) {

	assert := require.New(t)

	res := resource.Default()
	cfg := &OTelCollectorConfig{
		ServerURL:           "otelcollector.plivops.com:443",
		SecretToken:         "token",
		TransactionSampling: 0.5,
		Protocol:            "http",
	}
	tp, err := newOTelCollectorTracerProvider(context.TODO(), res, cfg)
	assert.Nil(err)
	assert.NotNil(tp)
}

func TestNewOTelCollectorTracerProviderFailure(t *testing.T) {

	assert := require.New(t)

	res := resource.Default()
	cfg := &OTelCollectorConfig{
		ServerURL:           "",
		SecretToken:         "token",
		TransactionSampling: 0.5,
		Protocol:            "http",
	}
	tp, err := newOTelCollectorTracerProvider(context.TODO(), res, cfg)
	assert.NotNil(err)
	assert.Nil(tp)
}

func TestNewOTelCollectorExporterGrpc(t *testing.T) {

	assert := require.New(t)

	var exporter *otlptrace.Exporter
	var err error

	cfg := &OTelCollectorConfig{
		ServerURL:           "otelcollector.plivops.com:443",
		SecretToken:         "token",
		TransactionSampling: 0.5,
		Protocol:            "grpc",
	}
	exporter, err = newOTelCollectorExporter(context.TODO(), cfg)
	assert.NotNil(exporter)
	assert.Nil(err)

	// should work for protocol case mismatch
	cfg.Protocol = "gRPC"
	exporter, err = newOTelCollectorExporter(context.TODO(), cfg)
	assert.NotNil(exporter)
	assert.Nil(err)
}
