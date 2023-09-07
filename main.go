package main

import "log"

const configFileName string = "pgexecute.yaml"

func main() {

	config, err := NewConfig(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("configuration: %#v", config)

}
