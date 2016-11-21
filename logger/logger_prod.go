// +build !debug

package logger

func (l *stdLogger) Debug(message string, context map[string]interface{}) {
}

func (l *GelfLogger) Debug(message string, context map[string]interface{}) {
}
