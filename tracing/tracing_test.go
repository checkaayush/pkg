package tracing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {

	assert := require.New(t)

	// nil config passed
	closer, err := Init(nil)
	assert.NotNil(closer)
	assert.Nil(err)

	// tracing not enabled
	closer, err = Init(&Config{
		Enabled: false,
	})
	assert.NotNil(closer)
	assert.Nil(err)

	// invalid provider passed
	closer, err = Init(&Config{
		Enabled:  true,
		Provider: "invalid",
	})
	assert.NotNil(closer)
	assert.NotNil(err)

	// empty provider passed
	closer, err = Init(&Config{
		Enabled:  true,
		Provider: "",
	})
	assert.NotNil(closer)
	assert.Nil(err)

	// test elastic provider
	closer, err = Init(&Config{
		Enabled:  true,
		Provider: "elastic",
		Elastic: &ElasticConfig{
			ServerURL:           "https://apm-server-73-usw1.non-prod.plivops.com:443",
			SecretToken:         "token",
			ServiceName:         "testsvc",
			ServiceVersion:      "version1",
			TransactionSampling: 0.5,
			Environment:         "dev",
		},
	})

	assert.NotNil(closer)
	assert.Nil(err)
}

func TestInitialize(t *testing.T) {

	assert := require.New(t)
	ctx := context.TODO()

	// nil config passed
	shutdown, err := Initialize(ctx, nil)
	assert.NotNil(err)
	assert.Nil(shutdown)

	// tracing not enabled
	shutdown, err = Initialize(ctx, &Config{
		Enabled: false,
	})
	assert.NotNil(shutdown)
	assert.Nil(err)

	// invalid provider passed
	shutdown, err = Initialize(ctx, &Config{
		Enabled:  true,
		Provider: "invalid",
	})
	assert.Nil(shutdown)
	assert.NotNil(err)

	// test noop provider
	shutdown, err = Initialize(ctx, &Config{
		Enabled:  true,
		Provider: "noop",
	})
	assert.NotNil(shutdown)
	assert.Nil(err)

	// test elastic provider
	shutdown, err = Initialize(ctx, &Config{
		Enabled:  true,
		Provider: "elastic",
		Elastic: &ElasticConfig{
			ServerURL:           "https://apm-server-73-usw1.non-prod.plivops.com:443",
			SecretToken:         "token",
			ServiceName:         "testsvc",
			ServiceVersion:      "version1",
			TransactionSampling: 0.5,
			Environment:         "dev",
			Protocol:            "http",
		},
	})
	assert.NotNil(shutdown)
	assert.Nil(err)

	// test jaeger provider
	shutdown, err = Initialize(ctx, &Config{
		Enabled:  true,
		Provider: "jaeger",
		Jaeger: &JaegerConfig{
			Host: "localhost",
			Port: 14268,
		},
	})
	assert.NotNil(shutdown)
	assert.Nil(err)

	// test otel collector provider (http)
	shutdown, err = Initialize(ctx, &Config{
		Enabled:  true,
		Provider: "otel_collector",
		OTelCollector: &OTelCollectorConfig{
			ServerURL:           "https://apm-server-73-usw1.non-prod.plivops.com:443",
			SecretToken:         "token",
			TransactionSampling: 0.5,
			Protocol:            "http",
		},
	})
	assert.NotNil(shutdown)
	assert.Nil(err)

	// test otel collector provider (grpc)
	shutdown, err = Initialize(ctx, &Config{
		Enabled:  true,
		Provider: "otel_collector",
		OTelCollector: &OTelCollectorConfig{
			ServerURL:           "apm-server-73-usw1.non-prod.plivops.com",
			SecretToken:         "token",
			TransactionSampling: 0.5,
			Protocol:            "grpc",
		},
	})
	assert.NotNil(shutdown)
	assert.Nil(err)

	// test mismatching provider and config
	shutdown, err = Initialize(ctx, &Config{
		Enabled:  true,
		Provider: "elastic",
		Jaeger: &JaegerConfig{
			Host: "localhost",
			Port: 14268,
		},
	})
	assert.Nil(shutdown)
	assert.NotNil(err)
}

func TestRemoveURLSchemePrefix(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{
				url: "",
			},
			want: "",
		},
		{
			name: "valid url with prefix",
			args: args{
				url: "https://example.com:443",
			},
			want: "example.com:443",
		},
		{
			name: "valid url without prefix",
			args: args{
				url: "example.com:443",
			},
			want: "example.com:443",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeURLSchemePrefix(tt.args.url); got != tt.want {
				t.Errorf("removeURLSchemePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
