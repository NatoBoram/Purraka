package main

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

func selectBestZScore() (selected selectedZScoreItem, err error) {
	err = db.QueryRow("SELECT "+
		"`data-wearableitemid`, `data-itemid`, `data-type`, `rarity-marker`, `abstract-name`, `abstract-type`, `currentPrice`, `zscore-currentPrice`, `buyNowPrice`, `zscore-buyNowPrice`, `data-bids`, `zscore-data-bids`, `abstract-icon` "+
		"FROM `z-market` "+
		"WHERE `data-type` != 'EggItem' "+
		"AND `data-bids` = 0 "+ // Block bids
		// "ORDER BY if(`data-bids` = 0, least(`zscore-buyNowPrice`, `zscore-currentPrice`), `zscore-currentPrice`) asc, if(`data-bids` = 0, greatest(`currentPrice`, `buyNowPrice`), `currentPrice`) desc "+
		"ORDER BY `zscore-buyNowPrice` asc, `buyNowPrice` desc "+
		"LIMIT 1;").Scan(&selected.datawearableitemid, &selected.dataitemid, &selected.datatype, &selected.raritymarker, &selected.abstractname, &selected.abstracttype, &selected.currentPrice, &selected.zscorecurrentPrice, &selected.buyNowPrice, &selected.zscorebuyNowPrice, &selected.databids, &selected.zscoredatabids, &selected.abstracticon)
	return selected, err
}

// Sent On Discord

func createSentOnDiscord() (res sql.Result, err error) {
	return db.Exec("create table if not exists `sent_on_discord` (`data-itemid` int PRIMARY KEY) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")
}

func insertSentOnDiscord(id int) (res sql.Result, err error) {
	return db.Exec("insert into `sent_on_discord` values(?);", id)
}

func selectSentOnDiscord(id int) (dataitemid int, err error) {
	db.QueryRow("select `data-itemid` from `sent_on_discord` where `data-itemid` = ?;", id).Scan(&dataitemid)
	return
}

// Callback Channel

func createCallbackChannel() (res sql.Result, err error) {
	return db.Exec("CREATE TABLE IF NOT EXISTS `callback_channel` (`guild` varchar(32) PRIMARY KEY, `channel` varchar(32)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")
}

func selectCallbackChannel(g *discordgo.Guild) (selected selectedGuildChannel, err error) {
	err = db.QueryRow("select `guild`, `channel` from `callback_channel` where `guild` = ?;", g.ID).Scan(&selected.guild, &selected.channel)
	return
}

func insertCallbackChannel(g *discordgo.Guild, c *discordgo.Channel) (res sql.Result, err error) {
	return db.Exec("insert into `callback_channel`(`guild`, `channel`) values(?, ?);", g.ID, c.ID)
}

func updateCallbackChannel(g *discordgo.Guild, c *discordgo.Channel) (res sql.Result, err error) {
	return db.Exec("update `callback_channel` set `channel` = ? where `guild` = ?;", c.ID, g.ID)
}

func deleteCallbackChannel(g *discordgo.Guild) (res sql.Result, err error) {
	return db.Exec("delete from `callback_channel` where `guild` = ?;", g.ID)
}
