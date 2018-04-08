package main

import (
	"database/sql"

	"github.com/NatoBoram/Purraka/config"
	"github.com/NatoBoram/Purraka/spider"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // http://go-database-sql.org/

func main() {

	// Load the configuration.
	err := config.Load()
	if err != nil {
		println("Couldn't load the config file.")
		println(err.Error())
		// config.Reset()
		return
	}

	// Open a connection to the database.
	db, err = sql.Open("mysql", config.DBConfig.User+":"+config.DBConfig.Password+"@tcp("+config.DBConfig.Address+":"+config.DBConfig.Port+")/"+config.DBConfig.Database)
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
	spider.Spider(db)
	go spider.StartSpider(db)

	// Wait for infinity
	<-make(chan struct{})
}
