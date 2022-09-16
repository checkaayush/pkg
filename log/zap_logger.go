package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	*zap.SugaredLogger
}

// With returns a new field logger which will log provided fields.
// Fields are specified as: f1, v1, f2, v2
func (l *zapLogger) With(args ...interface{}) Logger {
	return &zapLogger{
		l.SugaredLogger.With(args...),
	}
}

// WithAlert returns a new alert logger which will log the
// priority and type of alert.
func (l *zapLogger) WithAlert(alert Alert) Logger {
	return &zapLogger{
		l.SugaredLogger.With(
			"priority", alert.Priority(),
			"alert", alert.String()),
	}
}

// Write() will log the content provided at the Info level.
func (l *zapLogger) Write(p []byte) (n int, err error) {
	l.Infow(string(p))
	return len(p), nil
}

// NewZapLogger returns a new instance of *zap.SugaredLogger
// which implements the Logger interface.
func NewZapLogger(config *Config) (Logger, error) {

	var zapCfg zap.Config
	switch config.Environment {
	case "dev":
		fallthrough
	case "development":
		zapCfg = zap.NewDevelopmentConfig()
	case "prod":
		fallthrough
	case "production":
		zapCfg = zap.NewProductionConfig()
		zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	case "no-op":
		zapCfg = zap.NewProductionConfig()
		zapCfg.OutputPaths = []string{} // See https://pkg.go.dev/go.uber.org/zap#Open
	default:
		return nil, fmt.Errorf("invalid logger environment: valid values are dev/development/prod/production")
	}

	zl, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	logger := &zapLogger{
		zl.Sugar(),
	}

	logger.Sync()

	return logger, nil
}
