package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/bencicandrej/tricks/clock"
)

const (
	levelDebug = 0
	levelInfo  = 1
	levelError = 3
)

// Clock is the default implementation of something that returns
// the current time used for logging.
//
// This can be overriden in order to enable testing of the logger
// package using the clock.BrokenClock implementation.
var Clock = clock.New()

// Context is a shorthand for a map[string]interface.
type Context map[string]interface{}

// Logger is the simplest possible interface for a logger.
type Logger interface {
	Debug(message string, context map[string]interface{})
	Info(message string, context map[string]interface{})
	Error(message string, context map[string]interface{})
	Derive(prefix string) Logger
}

type stdLogger struct {
	output io.Writer
	prefix string
	*log.Logger
}

// NewStandardLogger initializes the stdLogger struct and returns
// it as a Logger interface implementation.
func NewStandardLogger(output io.Writer, prefix string) Logger {
	return &stdLogger{
		output: output,
		prefix: prefix,
	}
}

func (l *stdLogger) Info(message string, context map[string]interface{}) {
	l.log("INFO", message, context)
}

func (l *stdLogger) Error(message string, context map[string]interface{}) {
	l.log("ERROR", message, context)
}

func (l *stdLogger) log(level string, message string, context map[string]interface{}) {
	out := Clock.Now().Format("2006/01/02 15:04:05") + " [" + level + "] "
	if l.prefix != "" {
		out += l.prefix + " | "
	}

	out += message

	if context != nil {
		jsonContext, err := json.Marshal(context)
		if err == nil {
			out += fmt.Sprintf(" | %s", jsonContext)
		} else {
			out += fmt.Sprintf(` | {"marshalError": %q}`, err)
		}
	}

	if !strings.HasSuffix(out, "\n") {
		out = out + "\n"
	}

	fmt.Fprint(l.output, out)
}

func (l *stdLogger) Derive(prefix string) Logger {
	if l.Prefix() == "" {
		return NewStandardLogger(l.output, l.Prefix())
	}

	return NewStandardLogger(l.output, l.Prefix()+":"+prefix)
}

type GelfLogger struct {
	output io.Writer

	hostname string
	facility string
}

func NewGelfLogger(output io.Writer, hostname string, facility string) Logger {
	return &GelfLogger{
		output:   output,
		hostname: hostname,
		facility: facility,
	}
}

func (l *GelfLogger) Info(message string, context map[string]interface{}) {
	l.log(levelInfo, message, context)
}

func (l *GelfLogger) Error(message string, context map[string]interface{}) {
	l.log(levelError, message, context)
}

func (l *GelfLogger) Derive(facility string) Logger {
	if l.facility == "" {
		return NewGelfLogger(l.output, l.hostname, facility)
	}

	return NewGelfLogger(l.output, l.hostname, l.facility+":"+facility)
}

func (l *GelfLogger) log(level int, message string, context map[string]interface{}) {
	additionalFields := ""
	for key, value := range context {
		additionalFields += fmt.Sprintf(`,"_%s":"%v"`, key, value)
	}

	fmt.Fprintf(l.output, `{"version":"1.1","host":"%s","_facility":"%s","short_message":"%s","level":%d,"timestamp":%d%s}
`, l.hostname, l.facility, message, level, Clock.Now().Unix(), additionalFields)
}

type MultiLogger struct {
	loggers []Logger
}

func NewMultiLogger(loggers ...Logger) *MultiLogger {
	return &MultiLogger{
		loggers: loggers,
	}
}

func (l *MultiLogger) Add(logger Logger) {
	l.loggers = append(l.loggers, logger)
}

func (l *MultiLogger) Debug(message string, context map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Debug(message, context)
	}
}

func (l *MultiLogger) Info(message string, context map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Info(message, context)
	}
}

func (l *MultiLogger) Error(message string, context map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Error(message, context)
	}
}

func (l *MultiLogger) Derive(prefix string) Logger {
	newLoggers := []Logger{}
	for _, logger := range l.loggers {
		newLoggers = append(newLoggers, logger.Derive(prefix))
	}

	return NewMultiLogger(newLoggers...)
}
