package gologger

import (
	"errors"
	"testing"
)

func TestLogUDPMessageWithExtraFields(t *testing.T) {

	c := &LoggerConfig{
		GraylogHostname: "localhost",
		GraylogPort:     11111,
	}

	err := InitUDPLogger(c)

	if err != nil {
		t.Errorf("Error initializing logger %s", err.Error())
	}

	i := LogItems{
		ShortMsg:    "UDP Message with extra fields",
		FullMsg:     "full",
		ExtraFields: map[string]interface{}{"foo": "bar"},
		ErrorMsg:    nil,
	}

	err = GrayLogger.Info(i)

	if err != nil {
		t.Errorf("Error sending INFO message %s", err.Error())
	}
}

func TestLogUDPMessage(t *testing.T) {

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

	i := LogItems{
		ShortMsg:    "UDP Message",
		FullMsg:     "full",
		ExtraFields: map[string]interface{}{},
		ErrorMsg:    nil,
	}

	err = GrayLogger.Info(i)

	if err != nil {
		t.Errorf("Error sending INFO message %s", err.Error())
	}
}

func TestLogTCPMessage(t *testing.T) {

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

	i := LogItems{
		ShortMsg:    "TCP Message",
		FullMsg:     "full",
		ExtraFields: map[string]interface{}{},
		ErrorMsg:    nil,
	}

	err = GrayLogger.Info(i)

	if err != nil {
		t.Errorf("Error sending INFO message %s", err.Error())
	}
}

func TestLogTCPMessageWithExtraFields(t *testing.T) {

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

	i := LogItems{
		ShortMsg:    "TCP Message with extra fields",
		FullMsg:     "full",
		ExtraFields: map[string]interface{}{"foo": "bar"},
		ErrorMsg:    nil,
	}

	err = GrayLogger.Info(i)

	if err != nil {
		t.Errorf("Error sending INFO message %s", err.Error())
	}
}

func TestLogUDPErrorMessageWithExtraFields(t *testing.T) {
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

	var e = errors.New("Test error")

	i := LogItems{
		ShortMsg:    "UDP Message error message with extra fields",
		FullMsg:     "full",
		ExtraFields: map[string]interface{}{"foo": "bar"},
		ErrorMsg:    e,
	}

	err = GrayLogger.Error(i)

	if err != nil {
		t.Errorf("Error sending ERROR message %s", err.Error())
	}
}

func TestLogTCPErrorMessageWithExtraFields(t *testing.T) {
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

	var e = errors.New("Test error")

	i := LogItems{
		ShortMsg:    "TCP Message error message with extra fields",
		FullMsg:     "full",
		ExtraFields: map[string]interface{}{"foo": "bar"},
		ErrorMsg:    e,
	}

	err = GrayLogger.Error(i)

	if err != nil {
		t.Errorf("Error sending ERROR message %s", err.Error())
	}
}
