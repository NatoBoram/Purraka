package main

import "database/sql"

func createTables() (res sql.Result, err error) {

	// Declare tables to create
	functs := [...]func() (res sql.Result, err error){
		createSentOnDiscord,
		createCallbackChannel,
	}

	// Create the tables
	for _, funct := range functs {
		res, err = funct()
		if err != nil {
			return
		}
	}

	return
}
