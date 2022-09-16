package log

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoggerConfig(t *testing.T) {
	assert := require.New(t)

	// success cases
	tcs := []string{"dev", "development", "prod", "production", "no-op"}
	for _, tc := range tcs {
		l, err := NewLogger(&Config{
			Environment: tc,
		})
		assert.Nil(err)
		assert.NotNil(l)
	}

	// negative test cases
	tcs = []string{"", "DEV", "PROD", "invalid", "DEV", "PRODUCTION", "NOOP", "noop"}
	for _, tc := range tcs {
		l, err := NewLogger(&Config{
			Environment: tc,
		})
		assert.NotNil(err)
		assert.Nil(l)
	}
}
