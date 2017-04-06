package main

import (
	"encoding/json"
	"log"
	"os"
)

const (
	configPath = "./config.json"
)

type config struct {
	Source string
}

func (cfg *config) read() error {
	file, err := os.Open(configPath)
	if err != nil {
		log.Printf("Failed to open config: %v", err)
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Printf("Config file '%v' cannot be decoded: %v", file.Name(), err)
		return err
	}

	return nil
}

func (cfg config) write() error {
	file, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open config: %v", err)
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		log.Printf("Config file '%v' cannot be written: %v", file.Name(), err)
		return err
	}

	return nil
}
