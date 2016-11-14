package graylog

import (
	"fmt"
	"time"

	"github.com/bencicandrej/tricks/logger"

	"github.com/robertkowalski/graylog-golang"
)

// Logger is a graylog implementation of a logger.Logger interface.
type Logger struct {
	gelf *gelf.Gelf

	hostname  string
	namespace string
}

// WithHostname is a configuration function for setting a hostname for the
// Graylog GELF message.
func WithHostname(hostname string) func(*Logger) {
	return func(logger *Logger) {
		logger.hostname = hostname
	}
}

// WithNamespace is a configuration function for setting a namespace for the
// Graylog GELF message.
func WithNamespace(namespace string) func(*Logger) {
	return func(logger *Logger) {
		logger.namespace = namespace
	}
}

// SwitchNamespace is a shorthand for configuring a new logger with a different context.
func SwitchNamespace(l *Logger, namespace string) *Logger {
	newLogger := *l

	newLogger.namespace = namespace

	return &newLogger
}

// NewLogger returns a configured logger that
// sends data to a graylog instance.
func NewLogger(logger *gelf.Gelf, options ...func(*Logger)) logger.Logger {
	l := &Logger{
		gelf: logger,
	}

	for _, option := range options {
		option(l)
	}

	if l.hostname == "" {
		l.hostname = "localhost"
	}

	if l.namespace == "" {
		l.namespace = "default"
	}

	return l
}

// Info logs an info level message to graylog.
func (l *Logger) Info(message string, context map[string]interface{}) {
	l.log(1, message, context)
}

// Error log an error level message to graylog.
func (l *Logger) Error(message string, context map[string]interface{}) {
	l.log(3, message, context)
}

func (l *Logger) log(level int, message string, context map[string]interface{}) {
	additionalFields := ""
	for key, value := range context {
		additionalFields += fmt.Sprintf(`,"_%s": "%v"`, key, value)
	}

	l.gelf.Log(fmt.Sprintf(`
		"version": "1.1",
		"host": "%s",
		"_facility": "%s",
		"short_message": "%s",
		"level": %d,
		"timestamp": %d%s
	}`, l.hostname, l.namespace, message, level, time.Now().Unix(), additionalFields))

	fmt.Println(fmt.Sprintf(`
		"version": "1.1",
		"host": "%s",
		"_facility": "%s",
		"short_message": "%s",
		"level": %d,
		"timestamp": %d%s
	}`, l.hostname, l.namespace, message, level, time.Now().Unix(), additionalFields))
}
