package gologger

import "testing"

func TestInitLogger(t *testing.T) {
	c := LoggerConfig{
		GraylogHostname: "127.0.0.1",
		GraylogPort:     11111,
	}

	err := InitLogger(&c)

	if err != nil {
		t.Errorf("Error initializing logger %s", err.Error())
	}
}

func TestFailHostnameLogger(t *testing.T) {
	c := LoggerConfig{
		GraylogHostname: "fdandaibfa",
		GraylogPort:     11111,
	}

	err := InitLogger(&c)

	if err == nil {
		t.Errorf("Initialized invalid hostname logger")
		t.FailNow()
	}
}
