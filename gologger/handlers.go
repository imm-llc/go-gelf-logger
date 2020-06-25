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

// Warning logs a warning (level 4) message to Graylog
func Warning(s string, f string, ef map[string]interface{}) error {
	h, _ := os.Hostname()
	m := wrapBuildGraylogMessage(s, f, 4, ef, h)

	sendGraylogMessage(m)

	return nil
}

// Error logs an error (level 3) message to Graylog. This function requires you pass it an error that implements an Error() function
// {"ErrorMessage": e.Error()} will be appended to the extra fields passed to this function
func Error(s string, f string, ef map[string]interface{}, e error) error {
	h, _ := os.Hostname()
	ef["ErrorMessage"] = e.Error()

	m := wrapBuildGraylogMessage(s, f, 3, ef, h)

	sendGraylogMessage(m)

	return nil
}
