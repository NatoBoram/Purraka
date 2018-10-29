package main

import (
	"github.com/bwmarrin/discordgo"
)

func addHandlers(s *discordgo.Session) {

	s.AddHandler(messageHandler)

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Channel
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		printDiscordMessageError("Couldn't get a message's channel.", nil, c, m, err)
	}

	// Guild
	g, err := s.Guild(c.GuildID)
	if err != nil {
		printDiscordMessageError("Couldn't get a message's guild.", g, c, m, err)
	}

	// Ignore myself
	if m.Author.ID == me.ID {
		return
	}

	// Only isten to the owner, for now
	if m.Author.ID != g.OwnerID {
		return
	}

	// Check for a mention
	if len(m.Mentions) > 0 {

		// Mentionned me?
		if m.Mentions[0].ID == me.ID {
			commandHandler(s, g, c, m)
		}
	}
}
