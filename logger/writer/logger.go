package writer

import (
	"fmt"
	"io"
	"log"

	"github.com/bencicandrej/tricks/logger"
)

// Logger is a graylog implementation of a logger.Logger interface.
type Logger struct {
	l *log.Logger
}

// NewLogger returns a configured logger that
// sends data to a graylog instance.
func NewLogger(w io.Writer) logger.Logger {
	return &Logger{l: log.New(w, "", 0)}
}

// Info logs an info level message to graylog.
func (l *Logger) Info(message string, context map[string]interface{}) {
	l.log("INFO", message, context)
}

// Error log an error level message to graylog.
func (l *Logger) Error(message string, context map[string]interface{}) {
	l.log("ERROR", message, context)
}

func (l *Logger) log(level string, message string, context map[string]interface{}) {
	stringContext := ""
	for key, value := range context {
		stringContext += fmt.Sprintf(`,"_%s": "%v"`, key, value)
	}

	l.l.Printf("[ %s ] %s%s", level, message, stringContext)
}
