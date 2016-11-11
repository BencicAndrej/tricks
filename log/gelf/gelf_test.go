package gelf

import (
	"testing"
)

func TestGelfPacking(t *testing.T) {
	message := Message{
		Version:      "1.1",
		Host:         "localhost",
		ShortMessage: "Hello World",
	}

	expect := `{"version":"1.1","host":"localhost","short_message":"Hello World"}`

	packed := message.PackBuffer()
	if string(packed) != expect {
		t.Errorf("Message.Pack: got \n\t'%s'\n, wanted \n\t'%s'", packed, expect)
	}
}
