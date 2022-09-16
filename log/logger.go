package log

import "io"

// Refer: https://github.com/uber-go/zap/blob/master/FAQ.md#why-arent-logger-and-sugaredlogger-interfaces

// fieldLogger contains logger methods that support
// logging with fields (key-value pairs)
type fieldLogger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
}

// syncer contains Sync() method that flushes the
// internal buffer of logger.
type syncer interface {
	Sync() error
}

// Logger is an interface that all logging systems
// will implement.
type Logger interface {
	fieldLogger
	syncer
	// With returns a Logger containing fields provided
	// and is similar to logrus's FieldLogger
	With(args ...interface{}) Logger
	WithAlert(alert Alert) Logger
	io.Writer
}

var NewLogger = NewZapLogger
