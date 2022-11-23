# tracing

Package `tracing` provides tracing implementation compliant with OpenTelemetry specification.

Following providers are supported:

- **OpenTelemetry Collector** (HTTP & gRPC)
- **Jaeger** (HTTP only): For local testing
- **Noop**: Mock tracer provider for use in unit testing

## Local Development

- Bring up [Jaegar's all-in-one docker setup](https://www.jaegertracing.io/docs/1.37/getting-started/) locally.
- Use the following sample config for initialising tracing via Jaeger.

```go
cfg := &tracing.Config{
    Enabled:     true,
    Provider:    "jaeger",
    ServiceName: "<your-service-name>",
    Jaeger: &tracing.JaegerConfig{
        Host: "localhost",
        Port: 14268,
    },
}
```

## OpenTelemetry Usage

Initialize and set a global tracer provider:

```go
import (
        "github.com/plivo/pkg/tracing"
        "github.com/plivo/pkg/version"
)

func main() {
        ...

        ctx := context.Background()
        
        // Set the service version using pkg/version (recommended) to be used by OpenTelemetry.
        // You could also set it to any other desired version without using pkg/version.
        cfg.Tracing.ServiceVersion = version.String()

        // Global tracer provider is set
        shutdown, err := tracing.Initialize(ctx, cfg.Tracing)
        if err != nil {
                logger.Warnw("tracing.Initialize() failed", "error", err.Error())
        } else {
                // Handle shutdown properly so nothing leaks.
                defer func() { _ = shutdown(ctx) }()
        }

        ...
}
```

Insert spans wherever applicable:

```go

const tracerName = "redis"

func (s *RedisStore) SetMessageDetail(ctx context.Context, messageUUID string, detail *MessageDetail) error {

        ctx, span := otel.Tracer(tracerName).Start(ctx, "store.SetMessageDetail")
        defer span.End()

        ...

}
```

## OpenTracing Usage (Deprecated)

Initialize and set a global tracer:

```go
import "github.com/plivo/pkg/tracing"

func main() {
        ...

        if cfg.Tracing.Elastic != nil {
                cfg.Tracing.Elastic.ServiceVersion = version.String()
        }

        // Global opentracing tracer is set
        if err := tracing.Init(cfg.Tracing); err != nil {
                logger.Warnw("tracing.Init() failed", "error", err.Error())
        }

        ...
}
```

Insert spans wherever applicable:

```go
func (s *RedisStore) SetMessageDetail(ctx context.Context, messageUUID string, detail *MessageDetail) error {

        span, ctx := opentracing.StartSpanFromContext(ctx, "store.SetMessageDetail")
        defer span.Finish()

        ...

}
```

You can pass `context.Background()` to context at the place where handling a request originates.
