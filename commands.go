package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func commandHandler(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.MessageCreate) {

	// Split
	command := strings.Split(m.Content, " ")

	// Purraka
	if len(command) > 1 {
		switch command[1] {

		// Purraka get
		case "get":
			if len(command) > 2 {
				switch command[2] {

				// Purraka get channel
				case "channel":
					if len(command) > 3 {
						switch command[3] {

						// Purraka get channel callback
						case "callback":
							getCallbackChannelCommand(s, g, c, m)
						}
					}
				}
			}

		// Purraka set
		case "set":

			if len(command) > 2 {
				switch command[2] {

				// Purraka get channel
				case "channel":
					if len(command) > 3 {
						switch command[3] {

						// Purraka get channel callback
						case "callback":
							setCallbackChannelCommand(s, g, c, m)
						}
					}
				}
			}
		}
	}
}
