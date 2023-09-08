package main

import (
	"log"
)

const configFileName string = "pgexecute.yaml"

func main() {

	//	read configuration
	config, err := NewConfig(configFileName)
	if err != nil {
		log.Fatal(err)
	}

	//	execute scripts
	err = pgexecute(config.Connection, config.Files.In, config.Files.Out)
	if err != nil {
		log.Fatal(err)
	}
}
