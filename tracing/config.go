package tracing

type Config struct {
	Enabled        bool        `toml:"enabled" json:"enabled"`
	ServiceName    string      `toml:"service_name" validate:"required" json:"service_name"`
	ServiceVersion string      `toml:"service_version,omitempty" validate:"isdefault" json:"service_version"` // must not be in config
	Environment    string      `toml:"environment" validate:"required" json:"environment"`
	Provider       string      `toml:"provider" validate:"oneof=otlp noop" json:"provider"`
	OTLP           *OTLPConfig `toml:"otlp" json:"otlp"` // OTLP = OpenTelemetry Protocol
}

type OTLPConfig struct {
	ServerURL           string  `toml:"server_url" validate:"required" json:"server_url"`
	SecretToken         string  `toml:"secret_token" json:"secret_token"`
	Protocol            string  `toml:"protocol" validate:"oneof=http grpc" json:"protocol"`
	TransactionSampling float64 `toml:"transaction_sampling" validate:"required,gte=0,lte=1" json:"transaction_sampling"`
}
