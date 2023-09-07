package main

import (
	"log"
	"os"
)

const configFileName string = "pgexecute.yaml"

func main() {

	config, err := NewConfig(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("configuration: %#v", config)

	files, err := os.ReadDir(config.Files.In)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.Type().IsDir() {
			log.Printf("file: %v", file.Name())
		}
	}

}
