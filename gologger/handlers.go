package gologger

import (
	"os"
)

// Info logs an info (level 6) message to Graylog with the provided short message (s), full message (f), and extra fields (ef)
func (l LogItems) Info() error {
	h, _ := os.Hostname()
	m := wrapBuildGraylogMessage(l.ShortMsg, l.FullMsg, 6, l.ExtraFields, h)

	sendGraylogMessage(m)

	return nil
}

// Warning logs a warning (level 4) message to Graylog
func (l LogItems) Warning() error {
	h, _ := os.Hostname()
	m := wrapBuildGraylogMessage(l.ShortMsg, l.FullMsg, 4, l.ExtraFields, h)

	sendGraylogMessage(m)

	return nil
}

// Error logs an error (level 3) message to Graylog. This function requires you pass it an error that implements an Error() function
// {"ErrorMessage": e.Error()} will be appended to the extra fields passed to this function
func (l LogItems) Error() error {
	h, _ := os.Hostname()

	l.ExtraFields["ErrorMessage"] = l.ErrorMsg.Error()

	m := wrapBuildGraylogMessage(l.ShortMsg, l.FullMsg, 3, l.ExtraFields, h)

	sendGraylogMessage(m)

	return nil
}
