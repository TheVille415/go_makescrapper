package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type Lucy struct {
	Episode string
	Title   string
	Airdate string
}

func createJson(characters []Lucy) {
	jsonFile, _ := json.MarshalIndent(characters, "", " ")
	_ = ioutil.WriteFile("Lucy.json", jsonFile, 0644)
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	character := make([]Lucy, 0)

	// On every a element which has href attribute call callback
	c.OnHTML("table.general", func(e *colly.HTMLElement) {
		e.ForEach("tr.general-header", func(_ int, e *colly.HTMLElement) {
			newInfo := Lucy{}
			newInfo.Episode = e.ChildText("td")
			newInfo.Title = e.ChildText("a")
			newInfo.Airdate = e.ChildText("td")
			character = append(character, newInfo)
		})
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Received error:", e)
	})

	// Start scraping on https://lucifer.fandom.com/wiki/Episodes
	c.Visit("https://lucifer.fandom.com/wiki/Episodes")
}
