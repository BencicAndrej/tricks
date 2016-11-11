package log

import (
	"github.com/robertkowalski/graylog-golang"
)

type Context map[string]interface{}

type Logger interface {
	Info(message string, context Context)
	Error(message string, context Context)
}

type graylogLogger struct {
	logger *gelf.Gelf
}

// NewGraylogLogger returns a configured logger that
// sends data to a graylog instance.
func NewGraylogLogger(logger *gelf.Gelf) Logger {
	return &graylogLogger{
		logger: logger,
	}
}

func (l *graylogLogger) Info(message string, context Context) {

}

func (l *graylogLogger) Error(message string, context Context) {

}
