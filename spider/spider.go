package spider

import (
	"strconv"
	"time"
)

// StartSpider starts the spider. Call this function in a goroutine.
func StartSpider() {
	nextTime := time.Now().Truncate(time.Minute)
	nextTime = nextTime.Add(time.Minute)
	time.Sleep(time.Until(nextTime))
	Spider()
	go StartSpider()
}

// Spider scans Eldarya's market.
func Spider() {
	println("Spider : " + strconv.Itoa(time.Now().Minute()))
}
