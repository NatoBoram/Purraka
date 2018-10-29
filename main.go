package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// License
	fmt.Println("")
	fmt.Println("Purraka : Z-Scores from the market.")
	fmt.Println("Copyright © 2018 Nato Boram")
	fmt.Println("This program is free software : you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY ; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/.")
	fmt.Println("Contact : https://gitlab.com/NatoBoram/Purraka")
	fmt.Println("")

	// Database
	err := initDatabase()
	if err != nil {
		return
	}
	defer db.Close()

	// Discord
	session, err = initDiscord()
	if err != nil {
		return
	}
	defer session.Close()

	// Header
	err = readHeader(&header)
	if err != nil {
		return
	}

	// Open Discord Bot

	// Start the spider
	go StartSpider()
	go startCallback()

	// Wait for infinity
	<-make(chan struct{})
}

func initDatabase() (err error) {

	// Read the database config
	var database Database
	err = readDatabase(&database)
	if err != nil {
		fmt.Println("Could not load the database configuration.")
		fmt.Println(err.Error())
		writeTemplateDatabase()
		return
	}

	// Check for empty JSON
	if database.isEmpty() {
		err = errors.New("Configuration is missing inside " + databasePath)
		fmt.Println(err.Error())
		return
	}

	// Connect to MariaDB
	db, err = sql.Open("mysql", database.toConnectionString())
	if err != nil {
		fmt.Println("Could not connect to the database.")
		fmt.Println(err.Error())
		return
	}

	// Test the connection to MariaDB
	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping the database.")
		fmt.Println(err.Error())
		return
	}

	// Create the tables if they don't exist
	_, err = createTables()
	if err != nil {
		fmt.Println("Could not create a table in the database.")
		fmt.Println(err.Error())
		return
	}

	// Check version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	println("Connected to :", version)

	return
}

func initDiscord() (s *discordgo.Session, err error) {

	// Read the Discord config
	var discord Discord
	err = readDiscord(&discord)
	if err != nil {
		fmt.Println("Could not load the Discord configuration.")
		fmt.Println(err.Error())
		writeTemplateDiscord()
		return
	}

	// Check for empty JSON
	if discord.isEmpty() {
		err = errors.New("Configuration is missing inside " + discordPath)
		fmt.Println(err.Error())
		return
	}

	// Create a Discord session
	s, err = discordgo.New("Bot " + discord.Token)
	if err != nil {
		fmt.Println("Could not create a Discord session.")
		fmt.Println(err.Error())
		return
	}

	// Connect to Discord
	err = s.Open()
	if err != nil {
		fmt.Println("Could not connect to Discord.")
		fmt.Println(err.Error())
		return
	}

	// Myself
	me, err = s.User("@me")
	if err != nil {
		fmt.Println("Couldn't get myself.")
		fmt.Println(err.Error())
		return
	}

	// Master
	master, err = s.User(discord.MasterID)
	if err != nil {
		fmt.Println("Couldn't recognize my master.")
		fmt.Println(err.Error())
		return
	}

	// Handlers
	addHandlers(s)

	return
}
