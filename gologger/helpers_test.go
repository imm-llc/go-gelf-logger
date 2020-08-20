package gologger

import (
	"errors"
	"testing"
)

func TestUDPHelperMessages(t *testing.T) {
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

	LogInfo("short info", "full info message", "testing", map[string]interface{}{"test": "helpers", "proto": "udp"}, nil)
	LogWarning("short warning", "full warning message", "testing", map[string]interface{}{"test": "helpers", "proto": "udp"}, nil)
	LogError("short error", "full error message", "testing", map[string]interface{}{"test": "helpers", "proto": "udp"}, errors.New("Testing error"))
}

//
func TestTCPHelperMessages(t *testing.T) {
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

	LogInfo("short info", "full info message", "testing", map[string]interface{}{"test": "helpers", "proto": "tcp"}, nil)
	LogWarning("short warning", "full warning message", "testing", map[string]interface{}{"test": "helpers", "proto": "tcp"}, nil)
	LogError("short error", "full error message", "testing", map[string]interface{}{"test": "helpers", "proto": "tcp"}, errors.New("Testing error"))

}
