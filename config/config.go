package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// DBStruct hosts Purraka's configuration.
type DBStruct struct {
	User     string
	Password string
	Address  string
	Port     string
	Database string
}

// HeaderStruct is the header of the request sent to the URL.
type HeaderStruct struct {
	Accept         string
	AcceptEncoding string
	AcceptLanguage string
	Connection     string
	Cookie         string
	Host           string
	Referer        string
	UserAgent      string
	XRequestedWith string
}

const (
	dbConfigPath     = "./Purraka/db.json"
	headerConfigPath = "./Purraka/header.json"
)

var (

	// DBConfig of Purraka.
	DBConfig DBStruct

	// HeaderConfig of Purraka.
	HeaderConfig HeaderStruct
)

// Load the config file.
func Load() error {

	// Load Database Info
	println("Loading DB...")
	file, err := ioutil.ReadFile(dbConfigPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &DBConfig)
	if err != nil {
		return err
	}

	// Load Header Info
	println("Loading Header...")
	file, err = ioutil.ReadFile(headerConfigPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &HeaderConfig)
	if err != nil {
		return err
	}

	// Done!
	println("Loaded.")
	return err
}

// Save the current configuration
func Save() error {
	println("Saving...")
	json, err := json.Marshal(DBConfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dbConfigPath, json, os.FileMode(int(0777)))
	if err != nil {
		return err
	}
	println("Saved.")
	return err
}

// Reset the defaults and save.
func Reset() {
	println("Reset...")

	// Database Info
	var db DBStruct
	db.User = "root"
	db.Password = ""
	db.Address = "localhost"
	db.Port = "3306"
	db.Database = "eldarya"
	DBConfig = db
	Save()
}
