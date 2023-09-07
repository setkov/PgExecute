package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Files struct {
		In  string // input dir path
		Out string // output dir path
	}
}

func NewConfig(fileName string) (*Config, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
