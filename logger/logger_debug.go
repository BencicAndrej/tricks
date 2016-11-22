// +build debug

package logger

func (l *stdLogger) Debug(message string, context map[string]interface{}) {
	l.log("DEBUG", message, context)
}

func (l *GelfLogger) Debug(message string, context map[string]interface{}) {
	l.log(levelDebug, message, context)
}
