package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Load the configuration.
	err := load()
	if err != nil {
		println("Couldn't load the config file.")
		println(err.Error)
		reset()
	}

	// Open a connection to the database.
	db, err := sql.Open("mysql",
		config.user+":"+config.password+"@tcp("+config.address+":"+config.port+")/"+config.database)
	if err != nil {
		println(err.Error)
	}

	//

	defer db.Close()
}
