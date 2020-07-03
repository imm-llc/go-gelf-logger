package gologger

import (
	"testing"
)

func TestLogInfo(t *testing.T) {
	c := LoggerConfig{
		GraylogHostname: "localhost",
		GraylogPort:     11111,
	}

	var err error

	err = InitLogger(&c)

	if err != nil {
		t.Errorf("Error initializing logger %s", err.Error())
	}

	i := LogItems{
		ShortMsg:    "short",
		FullMsg:     "full",
		ExtraFields: map[string]interface{}{},
		ErrorMsg:    nil,
	}

	GrayLogger.Info(i)

	if err != nil {
		t.Errorf("Error sending INFO message %s", err.Error())
	}

}

func TestLogInfoWithExtraFields(t *testing.T) {
	c := LoggerConfig{
		GraylogHostname: "localhost",
		GraylogPort:     11111,
	}

	err := InitLogger(&c)

	if err != nil {
		t.Errorf("Error initializing logger %s", err.Error())
	}

	i := LogItems{
		ShortMsg:    "short",
		FullMsg:     "full",
		ExtraFields: map[string]interface{}{"foo": "bar"},
		ErrorMsg:    nil,
	}

	err = GrayLogger.Info(i)

	if err != nil {
		t.Errorf("Error sending INFO message %s", err.Error())
	}

}
