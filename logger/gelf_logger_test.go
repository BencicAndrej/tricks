package logger_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/bencicandrej/tricks/logger"
)

func TestGelfLoggerFormatting(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var out bytes.Buffer

			l := logger.NewGelfLogger(&out, "localhost", test.prefix)

			// Do only info for simplicity
			l.Info(test.message, test.context)

			outString := out.String()
			if outString != test.expectGelf {
				t.Errorf("gelfLogger.Info = %q, wanted %q", outString, test.expectGelf)
				return
			}
		})
	}
}

func TestGelfLoggerLevels(t *testing.T) {
	t.Run("Test INFO level logging", func(t *testing.T) {
		var out bytes.Buffer
		l := logger.NewGelfLogger(&out, "localhost", "")

		l.Info("Hello INFO Level Logging", nil)

		var outMap map[string]interface{}
		err := json.Unmarshal(out.Bytes(), &outMap)
		if err != nil {
			t.Errorf("error unmarshaling output: %v", err)
			return
		}

		if outMap["level"].(float64) != logger.LEVEL_INFO {
			t.Errorf("stdLogger.Info level = %d, wanted %d", outMap["level"].(float64), logger.LEVEL_INFO)
			return
		}
	})
	t.Run("Test ERROR level logging", func(t *testing.T) {
		var out bytes.Buffer
		l := logger.NewGelfLogger(&out, "localhost", "")

		l.Error("Hello ERROR Level Logging", nil)

		var outMap map[string]interface{}
		err := json.Unmarshal(out.Bytes(), &outMap)
		if err != nil {
			t.Errorf("error unmarshaling output: %v", err)
			return
		}

		if outMap["level"].(float64) != logger.LEVEL_ERROR {
			t.Errorf("stdLogger.Error level = %d, wanted %d", outMap["level"].(float64), logger.LEVEL_INFO)
			return
		}
	})
}

func TestGelfLoggerDerive(t *testing.T) {
	t.FailNow()
}
