package main

import "database/sql"

var (
	db *sql.DB // http://go-database-sql.org/

	// DBConfig of Purraka.
	DBConfig DBStruct

	// HeaderConfig of Purraka.
	HeaderConfig HeaderStruct
)
