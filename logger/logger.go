package logger

// Context is a shorthand for a map[string]interface.
type Context map[string]interface{}

// Logger is the simplest possible interface for a logger.
type Logger interface {
	Info(message string, context map[string]interface{})
	Error(message string, context map[string]interface{})
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

func (l *MultiLogger) Info(message string, context map[string]interface{})
