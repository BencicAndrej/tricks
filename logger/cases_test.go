package logger_test

import (
	"fmt"
	"os"
	"time"

	"github.com/bencicandrej/tricks/clock"
	"github.com/bencicandrej/tricks/logger"
)

func init() {
	t, err := time.Parse("2006-01-02", "2016-01-02")
	if err != nil {
		fmt.Printf("replace clock with a broken one: %v\n", err)
		os.Exit(1)
	}

	logger.Clock = clock.NewBrokenClock(t)
}

var tests = []struct {
	name       string
	message    string
	prefix     string
	context    map[string]interface{}
	expectStd  string
	expectGelf string
}{
	{
		name:      "Message only",
		message:   "Hello World",
		expectStd: "2016/01/02 00:00:00 [INFO] Hello World\n",
		expectGelf: `{"version":"1.1","host":"localhost","_facility":"","short_message":"Hello World","level":1,"timestamp":1451692800}
`,
	},
	{
		name:      "Message and prefix",
		message:   "Hello World",
		prefix:    "prefix",
		expectStd: "2016/01/02 00:00:00 [INFO] prefix | Hello World\n",
		expectGelf: `{"version":"1.1","host":"localhost","_facility":"prefix","short_message":"Hello World","level":1,"timestamp":1451692800}
`,
	},
	{
		name:    "Message, prefix and context",
		message: "Hello World",
		prefix:  "prefix",
		context: logger.Context{
			"key":    "value",
			"number": 3,
		},
		expectStd: "2016/01/02 00:00:00 [INFO] prefix | Hello World | {\"key\":\"value\",\"number\":3}\n",
		expectGelf: `{"version":"1.1","host":"localhost","_facility":"prefix","short_message":"Hello World","level":1,"timestamp":1451692800,"_key":"value","_number":"3"}
`,
	},
}
