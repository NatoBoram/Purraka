package main

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

var (
	db *sql.DB

	session *discordgo.Session
	me      *discordgo.User
	master  *discordgo.User

	header Header
)
