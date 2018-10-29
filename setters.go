package main

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

func setCallbackChannel(g *discordgo.Guild, c *discordgo.Channel) (res sql.Result, err error) {

	// Select
	_, err = selectCallbackChannel(g)
	if err == sql.ErrNoRows {

		// Insert
		res, err = insertCallbackChannel(g, c)
		if err != nil {
			return
		}

	} else if err == nil {

		// Update
		res, err = updateCallbackChannel(g, c)
		return

	}
	return
}
