package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func printDiscordMessageError(description string, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.MessageCreate, err error) {

	var log string

	if g != nil {
		log += "Guild : " + g.Name + "\n"
	}

	if c != nil {
		log += "Channel : " + c.Name + "\n"
	}

	if m != nil {
		log += "Author : " + m.Author.Username + "\n"
		log += "Message : " + m.Content + "\n"
	}

	if err != nil {
		log += "Error : " + err.Error() + "\n"
	}

	log += "\n"

	fmt.Println(log)
}
