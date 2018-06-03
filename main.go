package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Load the configuration.
	err := Load()
	if err != nil {
		println("Couldn't load the config file.")
		println(err.Error())
		// Reset()
		return
	}

	// Open a connection to the database.
	db, err = sql.Open("mysql", DBConfig.User+":"+DBConfig.Password+"@tcp("+DBConfig.Address+":"+DBConfig.Port+")/"+DBConfig.Database)
	if err != nil {
		println("Couldn't connect to the database.")
		println(err.Error())
		return
	}
	defer db.Close()

	// Check version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	println("Connected to :", version)

	// Open Discord Bot

	// Start the spider
	Spider(db)
	go StartSpider(db)

	// Wait for infinity
	<-make(chan struct{})
}
