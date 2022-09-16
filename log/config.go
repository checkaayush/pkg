package log

type Config struct {
	Environment string `json:"environment" toml:"environment" validate:"oneof=dev development prod production no-op"`
}
