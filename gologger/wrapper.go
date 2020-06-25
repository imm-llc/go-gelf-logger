package gologger

import (
	"log"
	"time"

	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

func wrapBuildGraylogMessage(shortMsg string, fullMsg string, level int32, extraFields map[string]interface{}, hostname string) *gelf.Message {

	m := &gelf.Message{
		Version:  "1.1",
		Host:     hostname,
		Short:    shortMsg,
		Full:     fullMsg,
		TimeUnix: float64(time.Now().Unix()),
		Level:    level,
		Extra:    extraFields,
	}

	return m
}

func sendGraylogMessage(m *gelf.Message) {
	err := gelfWriter.WriteMessage(m)

	if err != nil {
		log.Println("ERROR: Error sending message to Graylog", err.Error())
	}
}
