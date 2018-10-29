package main

import "github.com/bwmarrin/discordgo"

func setCallbackChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.MessageCreate) {

	// Set
	_, err := setCallbackChannel(g, c)
	if err != nil {
		printDiscordMessageError("Couldn't set a callback channel.", g, c, m, err)
	}

	// Typing
	err = s.ChannelTyping(c.ID)
	if err != nil {
		printDiscordMessageError("Couldn't tell I'm typing before setting my callback channel.", g, c, m, err)
	}

	// Send the channel
	_, err = s.ChannelMessageSend(c.ID, "The **callback** channel is now <#"+c.ID+">.")
	if err != nil {
		printDiscordMessageError("Couldn't send the callback channel.", g, c, m, err)
	}
}
