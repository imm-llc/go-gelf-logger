package gologger

import (
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

// LoggerConfig has the graylog server hostname (e.g. logs.example.com) and the input port (e.g. 12201)
type LoggerConfig struct {
	GraylogHostname string
	GraylogPort     int
}

// LogItems are the items sent in a log message
type LogItems struct {
	ShortMsg    string
	FullMsg     string
	ExtraFields map[string]interface{}
	ErrorMsg    error
}

// GrayLogger provides an interface to call our logging funcs
type GrayLogger interface {
	Info() error
	Warning() error
	Error() error
}

// Writer provides pointers to UDP and TCP Writers. One of the two will be nil
type Writer struct {
	UDPWriter *gelf.UDPWriter
	TCPWriter *gelf.TCPWriter
}
