package main

import (
	"github.com/bwmarrin/discordgo"
)

func getCallbackChannel(s *discordgo.Session, g *discordgo.Guild) (channel *discordgo.Channel, err error) {
	selected, err := selectCallbackChannel(g)
	if err != nil {
		return
	}
	return s.Channel(selected.channel)
}
