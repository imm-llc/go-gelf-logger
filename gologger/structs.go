package gologger

// LoggerConfig has the graylog server hostname (e.g. logs.example.com) and the input port (e.g. 12201)
type LoggerConfig struct {
	GraylogHostname string
	GraylogPort     int
}
