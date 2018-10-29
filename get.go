package main

import (
	"github.com/bwmarrin/discordgo"
)

func getCallbackChannelCommand(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.MessageCreate) {

	// Get the channel
	channel, err := getCallbackChannel(s, g)
	if err != nil {
		printDiscordMessageError("Couldn't get a channel.", g, c, m, err)
		return
	}

	// Typing
	err = s.ChannelTyping(c.ID)
	if err != nil {
		printDiscordMessageError("Couldn't tell I'm typing before sending my callback channel.", g, c, m, err)
	}

	// Send the channel
	_, err = s.ChannelMessageSend(c.ID, "The **callback** channel is <#"+channel.ID+">.")
	if err != nil {
		printDiscordMessageError("Couldn't send the callback channel.", g, c, m, err)
	}
}
