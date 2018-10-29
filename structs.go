package main

// Header is the header of the request sent to the URL.
type Header struct {
	Accept         string
	AcceptEncoding string
	AcceptLanguage string
	Connection     string
	Cookie         string
	Host           string
	Referer        string
	UserAgent      string
	XRequestedWith string
}

type item struct {
	id           string
	datatype     string
	icon         string
	rarity       string
	name         string
	abstracttype string
}

type sale struct {
	id           string
	itemid       string
	currentPrice string
	buyNowPrice  string
	bids         string
}

// Database hosts the bot's database configuration.
type Database struct {
	User     string
	Password string
	Address  string
	Port     int
	Database string
}

// Discord hosts the bot's Discord configuration.
type Discord struct {
	Token    string
	MasterID string
}

type selectedCallbackChannel struct {
	guild   string
	channel string
}

type selectedZScoreItem struct {
	datawearableitemid int
	dataitemid         int
	datatype           string
	raritymarker       string
	abstractname       string
	abstracttype       string
	currentPrice       int
	zscorecurrentPrice float64
	buyNowPrice        int
	zscorebuyNowPrice  float64
	databids           int
	zscoredatabids     float64
	abstracticon       string
}

type selectedGuildChannel struct {
	guild   string
	channel string
}
