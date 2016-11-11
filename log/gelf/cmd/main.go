package main

import (
	"fmt"

	"github.com/bencicandrej/tricks/log/gelf"
)

func main() {
	message := gelf.Message{
		Version:      "1.1",
		Host:         "localhost",
		ShortMessage: "Hello World",
	}

	expect := `{"version":"1.1","host":"localhost","short_message":"Hello World"}`

	packed := message.PackBuffer()
	if string(packed) != expect {
		fmt.Printf("Message.Pack: got \n\t'%s'\n, wanted \n\t'%s'", packed, expect)
	}
	fmt.Print(string(packed))
}
