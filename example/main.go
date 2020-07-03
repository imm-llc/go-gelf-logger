package main

import (
	"errors"

	"github.com/imm-llc/go-gelf-logger/example/subdir"
	bhiLogger "github.com/imm-llc/go-gelf-logger/gologger"
)

func main() {
	c := bhiLogger.LoggerConfig{
		GraylogHostname: "127.0.0.1",
		GraylogPort:     11111,
	}

	err := bhiLogger.InitLogger(&c)

	if err != nil {
		panic("Faied to set up logger")
	}

	subdir.LogMessage()
	//os.Exit(0)

	var i bhiLogger.LogItems

	i = bhiLogger.LogItems{
		ShortMsg:    "foo",
		FullMsg:     "foo bar whiz",
		ExtraFields: map[string]interface{}{"hello": "world", "x": 1},
		ErrorMsg:    nil,
	}

	bhiLogger.GrayLogger.Info(i)

	i = bhiLogger.LogItems{
		ShortMsg:    "foo",
		FullMsg:     "foo bar whiz",
		ExtraFields: map[string]interface{}{"hello": "world", "x": 1},
		ErrorMsg:    errors.New("Bad thing"),
	}

	bhiLogger.GrayLogger.Error(i)

}
