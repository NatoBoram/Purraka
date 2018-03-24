package main

import (
	"database/sql"

	"github.com/NatoBoram/Purraka/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Load the configuration.
	err := config.Load()
	if err != nil {
		println("Couldn't load the config file.")
		println(err.Error)
		config.Reset()
	}

	// Open a connection to the database.
	db, err := sql.Open("mysql",
		config.Config.user+":"+Config.password+"@tcp("+Config.address+":"+Config.port+")/"+Config.database)
	if err != nil {
		println(err.Error)
	}

	//

	defer db.Close()
}
