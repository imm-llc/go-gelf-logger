package main

import (
	bhiLogger "github.com/imm-llc/go-gelf-logger/gologger"
)

func main() {
	c := bhiLogger.LoggerConfig{
		GraylogHostname: "logs.juooce.com",
		GraylogPort:     11111,
	}

	err := bhiLogger.InitLogger(&c)

	if err != nil {
		panic("Faied to set up logger")
	}

	bhiLogger.Info("foo", "foo bar whiz", map[string]interface{}{"hello": "world", "x": 1})

}
