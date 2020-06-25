package gologger

import "testing"

func TestLogInfo(t *testing.T) {
	c := LoggerConfig{
		GraylogHostname: "logs.juooce.com",
		GraylogPort:     11111,
	}

	err := InitLogger(&c)

	if err != nil {
		t.Errorf("Error initializing logger %s", err.Error())
	}

	err = Info("short", "full", nil)

	if err != nil {
		t.Errorf("Error sending INFO message %s", err.Error())
	}

}

func TestLogInfoWithExtraFields(t *testing.T) {
	c := LoggerConfig{
		GraylogHostname: "logs.juooce.com",
		GraylogPort:     11111,
	}

	err := InitLogger(&c)

	if err != nil {
		t.Errorf("Error initializing logger %s", err.Error())
	}

	err = Info("short", "full", map[string]interface{}{"foo": "bar"})

	if err != nil {
		t.Errorf("Error sending INFO message %s", err.Error())
	}

}
