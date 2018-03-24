package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// StructConfig hosts Purraka's configuration.
type StructConfig struct {
	User     string
	Password string
	Address  string
	Port     string
	Database string
}

var (
	configPath = "./Purraka/config.json"

	// Configuration of Purraka.
	Configuration StructConfig
)

// Load the config file.
func Load() error {
	println("Loading...")
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &Configuration)
	if err != nil {
		return err
	}
	println("Loaded.")
	return nil
}

// Save the current configuration
func Save() error {
	println("Saving...")
	json, err := json.Marshal(Configuration)
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

// Reset the defaults and save.
func Reset() {
	println("Reset...")
	var defaults StructConfig
	defaults.User = "root"
	defaults.Password = ""
	defaults.Address = "localhost"
	defaults.Port = "3306"
	defaults.Database = "eldarya"
	Configuration = defaults
	Save()
}
