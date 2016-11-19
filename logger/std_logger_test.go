package logger_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bencicandrej/tricks/logger"
)

func TestStdLoggerFormatting(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var out bytes.Buffer

			l := logger.NewStandardLogger(&out, test.prefix)

			// Do only info for simplicity
			l.Info(test.message, test.context)

			cleanOutput := out.String()
			if cleanOutput != test.expectStd {
				t.Errorf("stdLogger.Info = %q, wanted %q", cleanOutput, test.expectStd)
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

func TestStdLoggerDerive(t *testing.T) {
	t.Error("Not implemented")
}
