package gologger

import (
	"testing"
)

func TestInitUDPLogger(t *testing.T) {
	c := &LoggerConfig{
		GraylogHostname: "localhost",
		GraylogPort:     11111,
	}

	err := InitUDPLogger(c)

	if err != nil {
		t.Errorf("Error initializing logger %s", err.Error())
	}

	if writer.UDPWriter == nil {
		t.Errorf("UDPWriter is nil")
	}
}

func TestInitTCPLogger(t *testing.T) {
	c := &LoggerConfig{
		GraylogHostname: "localhost",
		GraylogPort:     11111,
	}

	err := InitTCPLogger(c)

	if err != nil {
		t.Errorf("Error initializing logger %s", err.Error())
	}

	if writer.TCPWriter == nil {
		t.Errorf("TCPWriter is nil")
	}
}
