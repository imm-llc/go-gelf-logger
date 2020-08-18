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

func sendGraylogMessageV2(m *gelf.Message) {
	if writer.TCPWriter != nil {
		if err := writer.TCPWriter.WriteMessage(m); err != nil {
			log.Println("ERROR: Error sending message to Graylog", err.Error())
		}
	} else if writer.UDPWriter != nil {
		if err := writer.UDPWriter.WriteMessage(m); err != nil {
			log.Println("ERROR: Error sending message to Graylog", err.Error())
		}
	} else {

	}
}
