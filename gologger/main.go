package gologger

import (
	"fmt"
	"log"

	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

// InitLogger creates our gelf.UDPWriter from the provided LoggerConfig pointer
func InitLogger(c *LoggerConfig) error {
	var err error

	s := fmt.Sprintf("%s:%d", c.GraylogHostname, c.GraylogPort)
	gelfWriter, err = gelf.NewUDPWriter(s)

	if err != nil {
		log.Println("Error creating UDP writer", err.Error())
		return err
	}

	return nil

}
