package spider

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/NatoBoram/Purraka/config"
	"golang.org/x/net/html"
)

var url = "http://www.eldarya.fr/marketplace/ajax_search?from=0&to=10000"

// StartSpider starts the spider. Call this function in a goroutine.
func StartSpider(db *sql.DB) {
	for {
		nextTime := time.Now().Truncate(time.Minute)
		nextTime = nextTime.Add(time.Minute)
		time.Sleep(time.Until(nextTime))

		// Launch the spider
		err := Spider(db)
		if err != nil {
			println(err.Error())
		}
	}
}

// Spider scans Eldarya's market.
func Spider(db *sql.DB) error {

	// Start
	start := time.Now()
	println("Start :", strconv.Itoa(start.Hour())+":"+strconv.Itoa(start.Minute()))

	// Begin
	tx, err := db.Begin()
	if err != nil {
		println("Couldn't begin a transaction.")
		return err
	}

	// Query Item
	selectItem, err := tx.Prepare("select `abstract-name` from items where `data-wearableitemid` = ?;")
	if err != nil {
		println("Couldn't prepare the statement select item.")
		return err
	}
	defer selectItem.Close()

	// Prepare Item
	insertItem, err := tx.Prepare("insert into items(`data-wearableitemid`, `data-type`, `abstract-icon`, `rarity-marker`, `abstract-name`, `abstract-type`) values(?, ?, ?, ?, ?, ?);")
	if err != nil {
		println("Couldn't prepare the statement insert item.")
		return err
	}
	defer insertItem.Close()

	// Query Market
	selectSale, err := tx.Prepare("select `data-itemid` from market where `data-itemid` = ?;")
	if err != nil {
		println("Couldn't prepare the statement select sale.")
		return err
	}
	defer selectSale.Close()

	// Prepare Insert
	insertSale, err := tx.Prepare("insert into market(`data-itemid`, `data-wearableitemid`, `currentPrice`, `buyNowPrice`, `data-bids`) values(?, ?, ?, ?, ?);")
	if err != nil {
		println("Couldn't prepare the statement insert sale.")
		return err
	}
	defer insertSale.Close()

	// Prepare Update
	updateSale, err := tx.Prepare("update `market` set `currentPrice` = ?, `data-bids` = ?, `active` = 1 where `data-itemid` = ?;")
	if err != nil {
		println("Couldn't prepare the statement update sale.")
		return err
	}
	defer updateSale.Close()

	// Log
	println("Disabling every sales...")

	// Disable everything
	_, err = tx.Exec("update `market` set `active` = 0;")
	if err != nil {
		println("Couldn't disable every sales.")
		println(err.Error())
	}

	// Client
	client := &http.Client{}

	ok := true
	for page := 0; ok; page += 7 {

		var url = "http://www.eldarya.fr/marketplace/ajax_search?from=" + strconv.Itoa(page) + "&to=" + strconv.Itoa(page+6)

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

	L:
		for {
			tt := z.Next()
			switch tt {
			case html.ErrorToken:
				break L
			case html.StartTagToken:
				t := z.Token()

				switch t.Data {
				case "li": // New item

					// Item
					wearableitemid := tag(t, "data-wearableitemid")
					datatype := tag(t, "data-type")

					// Sale
					itemid := tag(t, "data-itemid")

					// Objects
					i = item{id: wearableitemid, datatype: datatype}
					s = sale{id: itemid, itemid: wearableitemid, currentPrice: "0", buyNowPrice: "0"}
					break
				case "img": // abstract-icon
					if tag(t, "class") == "abstract-icon" {
						i.icon = tag(t, "src")
					}
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
							s.currentPrice = tag(t, "data-price")
							s.bids = tag(t, "data-bids")
						} else {
							// This is a Current Price
							s.buyNowPrice = tag(t, "data-price")
						}
					}
					break
				}
				break
			case html.TextToken:
				t := z.Token()
				if i.name == "abstract-name" {
					i.name = strings.TrimSpace(t.Data)
				}
				if i.abstracttype == "abstract-type" {
					i.abstracttype = strings.TrimSpace(t.Data)
				}
				break
			case html.EndTagToken:
				t := z.Token()
				if t.Data == "li" {
					// Save in the list
					items = append(items, i)
					sales = append(sales, s)

					// Delete old values
					i = item{}
					s = sale{}
				}
				break
			}
		}

		// Log
		//println("There is", strconv.Itoa(len(items)), "items on this page.")
		if len(items) == 0 {
			ok = false
		}

		// Log
		//println("Going trough items...")

		// Insert Items
		for _, itemval := range items {
			var itemname string
			err := selectItem.QueryRow(itemval.id).Scan(&itemname)
			if err != nil {

				// New item
				println("New item :", itemval.name)
				_, err := insertItem.Exec(itemval.id, itemval.datatype, itemval.icon, itemval.rarity, itemval.name, itemval.abstracttype)
				if err != nil {
					println("Couldn't insert", itemval.name+".")
					println(err.Error())
				}
			}
		}

		// Log
		//println("Going trough sales...")

		// Insert Market
		for _, saleval := range sales {
			var saleid string
			err := selectSale.QueryRow(saleval.id).Scan(&saleid)
			if err != nil {

				// New sale
				println("New sale :", saleval.id)
				_, err := insertSale.Exec(saleval.id, saleval.itemid, saleval.currentPrice, saleval.buyNowPrice, saleval.bids)
				if err != nil {
					println("Couldn't insert", saleval.id+".")
					println(err.Error())
				}
			} else {

				// Old sale
				_, err := updateSale.Exec(saleval.currentPrice, saleval.bids, saleval.id)
				if err != nil {
					println("Couldn't update", saleval.id+".")
					println(err.Error())
				}
			}
		}
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		println("Couldn't commit the transaction.")
		println(err.Error())
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
}

func tag(token html.Token, key string) string {
	for _, attr := range token.Attr {
		if attr.Key == key {
			return strings.TrimSpace(attr.Val)
		}
	}
	return ""
}
