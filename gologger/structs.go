package gologger

// LoggerConfig has the graylog server hostname (e.g. logs.example.com) and the input port (e.g. 12201)
type LoggerConfig struct {
	GraylogHostname string
	GraylogPort     int
}

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
