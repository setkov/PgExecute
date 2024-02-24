package main

import (
	"log"
)

func main() {
	log.Print("Copyright: https://github.com/setkov/PgExecute (ver: 1.00)")

	//	read configuration
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//	execute scripts
	err = pgexecute(config.Connection, config.Files.In, config.Files.Out)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
