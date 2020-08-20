package gologger

import (
	"errors"
	"testing"
)

func TestSendHelperLogs(t *testing.T) {
	c := LoggerConfig{
		GraylogHostname: "localhost",
		GraylogPort:     11111,
	}

	err := InitLogger(&c)

	if err != nil {
		t.Errorf("Error creating logger")
		t.FailNow()
	}

	LogInfo("short info", "full info message", "testing", map[string]interface{}{"foo": "bar"}, nil)

	LogWarning("short warning", "full warning message", "testing", map[string]interface{}{"foo": "bar"}, nil)

	LogError("short error", "full error message", "testing", map[string]interface{}{"foo": "bar"}, errors.New("Testing error"))
}
