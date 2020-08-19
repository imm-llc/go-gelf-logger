package gologger

import (
	"fmt"
	"log"

	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

// InitUDPLogger initializes a new GELF UDP Logger
func InitUDPLogger(c *LoggerConfig) error {

	var err error
	writer = Writer{}

	writer.TCPWriter = nil

	s := fmt.Sprintf("%s:%d", c.GraylogHostname, c.GraylogPort)
	writer.UDPWriter, err = gelf.NewUDPWriter(s)

	if err != nil {
		log.Println("Error creating UDP writer", err.Error())
		return err
	}
	log.Println()

	return nil
}

// InitTCPLogger initializes a new GELF TCP Logger
func InitTCPLogger(c *LoggerConfig) error {

	var err error
	writer = Writer{}
	writer.UDPWriter = nil

	s := fmt.Sprintf("%s:%d", c.GraylogHostname, c.GraylogPort)
	writer.TCPWriter, err = gelf.NewTCPWriter(s)

	if err != nil {
		log.Println("Error creating TCP writer", err.Error())
		return err
	}

	return nil
}
