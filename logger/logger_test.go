package logger_test

import (
	"bytes"
	"fmt"
	"github.com/bencicandrej/tricks/logger"
	"strings"
	"testing"
)

func TestStdLoggerFormatting(t *testing.T) {
	tests := []struct {
		name    string
		message string
		prefix  string
		context map[string]interface{}
		expect  string
	}{
		{
			name:    "Message only",
			message: "Hello World",
			expect:  "[INFO] Hello World\n",
		},
		{
			name:    "Message and prefix",
			message: "Hello World",
			prefix:  "prefix",
			expect:  "[INFO] prefix | Hello World\n",
		},
		{
			name:    "Message, prefix and context",
			message: "Hello World",
			prefix:  "prefix",
			context: logger.Context{
				"key":    "value",
				"number": 3,
			},
			expect: "[INFO] prefix | Hello World | {\"key\":\"value\",\"number\":3}\n",
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("Test case #%d", index+1), func(t *testing.T) {
			var out bytes.Buffer

			l := logger.NewStandardLogger(&out, test.prefix)

			// Do only info for simplicity
			l.Info(test.message, test.context)

			cleanOutput := out.String()[20:]
			if cleanOutput != test.expect {
				t.Errorf("stdLogger.Info = %q, wanted %q", cleanOutput, test.expect)
				return
			}
		})
	}
}

func TestStdLoggerLevels(t *testing.T) {
	t.Run("Test INFO level logging", func(t *testing.T) {
		var out bytes.Buffer
		l := logger.NewStandardLogger(&out, "")

		l.Info("Hello INFO Level Logging", nil)

		cleanOutput := out.String()[20:]
		if !strings.HasPrefix(cleanOutput, "[INFO] ") {
			t.Errorf("stdLogger.Info = %q, wanted %q", cleanOutput, "[INFO] Hello INFO Level Logging\n")
			return
		}
	})
	t.Run("Test ERROR level logging", func(t *testing.T) {
		var out bytes.Buffer
		l := logger.NewStandardLogger(&out, "")

		l.Error("Hello ERROR Level Logging", nil)

		cleanOutput := out.String()[20:]
		if !strings.HasPrefix(cleanOutput, "[ERROR] ") {
			t.Errorf("stdLogger.Error = %q, wanted %q", cleanOutput, "[ERROR] Hello ERROR Level Logging\n")
			return
		}
	})
}
