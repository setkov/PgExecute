package main

import (
	"log"
)

func main() {

	//	read configuration
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	//	execute scripts
	err = pgexecute(config.Connection, config.Files.In, config.Files.Out)
	if err != nil {
		log.Fatal(err)
	}
}
