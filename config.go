package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Structure that hosts Purreko's configuration.
type configStruct struct {
	user     string
	password string
	address  string
	port     string
	database string
}

var (
	configPath = "./purreko/config.json"
	config     configStruct
)

func load() error {
	println("Loading...")
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}
	println("Loaded.")
	return nil
}

func save() error {
	println("Saving...")
	json, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(configPath, json, os.FileMode(int(0777)))
	if err != nil {
		return err
	}
	println("Saved.")
	return nil
}

func reset() {
	println("Reset...")
	var defaults configStruct
	defaults.user = "root"
	defaults.password = ""
	defaults.address = "localhost"
	defaults.port = "3306"
	defaults.database = "eldarya"
	config = defaults
	save()
}
