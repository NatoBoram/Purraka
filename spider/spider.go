package spider

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/NatoBoram/Purraka/config"
	"golang.org/x/net/html"
)

var url = "http://www.eldarya.fr/marketplace/ajax_search?from=0&to=10000"

// StartSpider starts the spider. Call this function in a goroutine.
func StartSpider() {
	for {
		nextTime := time.Now().Truncate(time.Minute)
		nextTime = nextTime.Add(time.Minute)
		time.Sleep(time.Until(nextTime))

		// Launch the spider
		err := Spider()
		if err != nil {
			println(err.Error())
		}
	}
}

// Spider scans Eldarya's market.
func Spider() error {

	// Start
	start := time.Now()
	println("Start :", strconv.Itoa(start.Hour())+":"+strconv.Itoa(start.Minute()))

	// Client
	client := &http.Client{}

	// Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println("Couldn't create a request!")
		return err
	}

	// Header
	req.Header.Add("Cookie", config.HeaderConfig.Cookie)

	// Response
	resp, err := client.Do(req)
	if err != nil {
		println("Couldn't receive a response!")
		return err
	}
	defer resp.Body.Close()

	// Tokenizer
	z := html.NewTokenizer(resp.Body)

	items := []item{}
	var i item

	sales := []sale{}
	var s sale

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			break
		case html.StartTagToken:
			t := z.Token()

			switch t.Data {
			case "li": // New item

				// Delete old values
				i = item{}
				s = sale{}

				// Item
				wearableitemid := tag(t, "data-wearableitemid")
				datatype := tag(t, "data-type")

				// Sale
				itemid := tag(t, "data-itemid")

				// Objects
				i = item{id: wearableitemid, datatype: datatype}
				s = sale{id: itemid}
				break
			case "img": // abstract-icon
				i.icon = tag(t, "src")
				break
			case "div":

				switch tag(t, "class") {

				// Rarity Marker
				case "rarity-marker-common":
					i.rarity = "common"
					break
				case "rarity-marker-rare":
					i.rarity = "rare"
					break
				case "rarity-marker-epic":
					i.rarity = "epic"
					break
				case "rarity-marker-legendary":
					i.rarity = "legendary"
					break
				case "rarity-marker-event":
					i.rarity = "event"
					break

					// Name
				case "abstract-name":
					i.name = "abstract-name"
					break

					// Type
				case "abstract-type":
					i.abstracttype = "abstract-type"
					break
				}

			case "span":

				if tag(t, "data-price") != "" {
					// This is a price.
					if tag(t, "data-bids") != "" {
						// This is a Buy Now Price
						s.buyNowPrice = tag(t, "data-price")
						s.bids = tag(t, "data-bids")
					} else {
						// This is a Current Price
						s.currentPrice = tag(t, "data-price")
					}
				}
				break
			}
			break
		case html.TextToken:
			t := z.Token()
			if i.name == "abstract-name" {
				i.name = strings.TrimSpace(t.Data)
				println("abstract-name :", i.name)
			}
			if i.abstracttype == "abstract-type" {
				i.abstracttype = strings.TrimSpace(t.Data)
				println("abstract-type :", i.abstracttype)
			}
			break
		case html.EndTagToken:
			t := z.Token()
			if t.Data == "li" {
				items = append(items, i)
				sales = append(sales, s)
			}
		}
	}

	// End
	end := time.Since(start)
	println("End :", end.String())

	return nil
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
	active       string
}

func tag(token html.Token, key string) string {
	for _, attr := range token.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
