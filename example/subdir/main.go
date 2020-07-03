package subdir

import (
	bhiLogger "github.com/imm-llc/go-gelf-logger/gologger"
)

func LogMessage() {

	i := bhiLogger.LogItems{
		ShortMsg:    "foo",
		FullMsg:     "foo bar whiz",
		ExtraFields: map[string]interface{}{"hello": "world", "x": 1},
		ErrorMsg:    nil,
	}

	bhiLogger.GrayLogger.Info(i)

}