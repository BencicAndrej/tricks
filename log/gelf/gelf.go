package gelf

import (
	"bytes"
	"strconv"
)

// Message represents a single unpacked GELF message.
type Message struct {
	Version          string
	Host             string
	ShortMessage     string
	FullMessage      string
	Timestamp        int
	Level            int
	AdditionalFields map[string]string
}

// PackBuffer formats the message into GELF format.
func (m *Message) PackBuffer() []byte {
	var buf bytes.Buffer

	buf.WriteString(`{"version":"`)
	buf.WriteString(m.Version)
	buf.WriteString(`","host":"`)
	buf.WriteString(m.Host)
	buf.WriteString(`","short_message":"`)
	buf.WriteString(m.ShortMessage)

	if m.FullMessage != "" {
		buf.WriteString(`","full_message":"`)
		buf.WriteString(m.FullMessage)
	}

	if m.Timestamp != 0 {
		buf.WriteString(`","timestamp":"`)
		buf.WriteString(strconv.Itoa(m.Timestamp))
	}

	if m.Level != 0 {
		buf.WriteString(`","level":"`)
		buf.WriteString(strconv.Itoa(m.Level))
	}

	buf.WriteString(`"`)

	for key, value := range m.AdditionalFields {
		buf.WriteString(`,"`)
		buf.WriteString(key)
		buf.WriteString(`":"`)
		buf.WriteString(value)
		buf.WriteString(`"`)
	}

	buf.WriteString(`}`)

	return buf.Bytes()
}

// PackConcat formats the message into GELF format.
func (m *Message) PackConcat() []byte {

	return nil
}
