package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func startCallback() {
	for {

		// Wait for next minute's cycle
		time.Sleep(time.Until(time.Now().Truncate(time.Minute).Add(time.Minute)))

		// Launch the callback
		callback()
	}
}

func callback() {

	// Select an item
	item, err := selectBestZScore()
	if err != nil {
		fmt.Println("Couldn't select the best ZScore item!")
		fmt.Println(err.Error())
		return
	}

	// Log it
	fmt.Println("\n" + "Selected " + item.abstractname)

	// Check if it's already been sent
	dataitemid, err := selectSentOnDiscord(item.dataitemid)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Couldn't check if this item already exists!")
		fmt.Println(err.Error())
		return
	}

	// Check
	if dataitemid == item.dataitemid {
		fmt.Println("This item has already been sent to Discord.")
		return
	}

	// Check!
	if session.State.Guilds == nil {
		return
	}

	for _, guild := range session.State.Guilds {

		// Get the callback channel
		channel, err := getCallbackChannel(session, guild)
		if err != nil {
			fmt.Println("This guild doesn't have a callback channel.")
			fmt.Println(err.Error())
			continue
		}

		// Create an embed
		embed := &discordgo.MessageEmbed{
			URL:   "https://eldarya.fr/marketplace",
			Title: item.abstractname,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://eldarya.fr" + item.abstracticon,
			},
			Fields: []*discordgo.MessageEmbedField{},
		}

		// Color
		color := colorFromRarity(item.raritymarker)
		if color != 0 {
			embed.Color = color
		}

		// Nom
		if item.abstractname != "" {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Nom",
				Value:  item.abstractname,
				Inline: true,
			})
		}

		// Catégorie
		if item.datatype != "" {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Catégorie",
				Value:  item.datatype,
				Inline: true,
			})
		}

		// Type
		if item.abstracttype != "" {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Type",
				Value:  item.abstracttype,
				Inline: true,
			})
		}

		// Rareté
		if item.raritymarker != "" {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Rareté",
				Value:  item.raritymarker,
				Inline: true,
			})
		}

		// Mise actuelle
		if item.currentPrice > 0 {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Mise actuelle",
				Value:  strconv.Itoa(item.currentPrice),
				Inline: true,
			})
		}

		// Cote Z de la mise actuelle
		if item.currentPrice > 0 {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Cote Z actuelle",
				Value:  strconv.FormatFloat(item.zscorecurrentPrice, 'f', 2, 64),
				Inline: true,
			})
		}

		// Achat immédiat
		if item.buyNowPrice > 0 {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Achat immédiat",
				Value:  strconv.Itoa(item.buyNowPrice),
				Inline: true,
			})
		}

		// Cote Z de l'achat immédiat
		if item.buyNowPrice > 0 {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Cote Z immédiate",
				Value:  strconv.FormatFloat(item.zscorebuyNowPrice, 'f', 2, 64),
				Inline: true,
			})
		}

		// Enchères
		if item.databids > 0 {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Enchères",
				Value:  strconv.Itoa(item.databids),
				Inline: true,
			})
		}

		// Cote Z du nombre d'enchères
		if item.databids > 0 {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Cote Z enchères",
				Value:  strconv.FormatFloat(item.zscoredatabids, 'f', 2, 64),
				Inline: true,
			})
		}

		// Send the item!
		_, err = session.ChannelMessageSendEmbed(channel.ID, embed)
		if err != nil {
			fmt.Println("Couldn't send an item.")
			fmt.Println(err.Error())
			continue
		}
	}

	// It's sent
	_, err = insertSentOnDiscord(item.dataitemid)
	if err != nil {
		fmt.Println("Couldn't save a sent item.")
		fmt.Println(err.Error())
	}

	return
}

func colorFromRarity(rarity string) (color int) {
	switch rarity {
	case "common":
		return
	case "rare":
		return 0x67e8f3
	case "epic":
		return 0xE485F5
	case "legendary":
		return 0xECC600
	case "event":
		return 0x4cec44
	}
	return
}
