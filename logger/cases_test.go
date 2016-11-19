package logger_test

import "github.com/bencicandrej/tricks/logger"

var tests = []struct {
	name       string
	message    string
	prefix     string
	context    map[string]interface{}
	expectStd  string
	expectGelf string
}{
	{
		name:       "Message only",
		message:    "Hello World",
		expectStd:  "[INFO] Hello World\n",
		expectGelf: "",
	},
	{
		name:       "Message and prefix",
		message:    "Hello World",
		prefix:     "prefix",
		expectStd:  "XXXXXXXXXXXXXXXXXX [INFO] prefix | Hello World\n",
		expectGelf: "",
	},
	{
		name:    "Message, prefix and context",
		message: "Hello World",
		prefix:  "prefix",
		context: logger.Context{
			"key":    "value",
			"number": 3,
		},
		expectStd:  "[INFO] prefix | Hello World | {\"key\":\"value\",\"number\":3}\n",
		expectGelf: "",
	},
}
