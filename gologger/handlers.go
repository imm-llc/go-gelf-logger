package gologger

import (
	"os"
)

// Info logs an info (level 6) message to Graylog with the provided short message (s), full message (f), and extra fields (ef)
func Info(s string, f string, ef map[string]interface{}) error {
	h, _ := os.Hostname()
	m := wrapBuildGraylogMessage(s, f, 6, ef, h)

	sendGraylogMessage(m)

	return nil
}
