package main

import (
	"fmt"
	//"regexp"

	"github.com/gocolly/colly/v2"
)

func scrape(search string) {
	c := colly.NewCollector(
		// Add Url rules here
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(search)
}