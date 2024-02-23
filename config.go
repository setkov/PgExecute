package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Connection string // postgresql connection string
	Files      struct {
		In  string // input dir path
		Out string // output dir path
	}
}

func NewConfig() (*Config, error) {
	fileName := "pgexecute.yaml"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	log.Printf("Read configuration from file: %v", fileName)

	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	log.Printf("Configuration: %v", config)
	return &config, nil
}
