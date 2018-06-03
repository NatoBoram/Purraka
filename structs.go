package main

type (
	// DBStruct hosts Purraka's configuration.
	DBStruct struct {
		User     string
		Password string
		Address  string
		Port     string
		Database string
	}

	// HeaderStruct is the header of the request sent to the URL.
	HeaderStruct struct {
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

	item struct {
		id           string
		datatype     string
		icon         string
		rarity       string
		name         string
		abstracttype string
	}

	sale struct {
		id           string
		itemid       string
		currentPrice string
		buyNowPrice  string
		bids         string
	}
)
