package main

import (
    "fmt"
    "os"
	"strings"

	"github.com/rocketlaunchr/google-search"
	"github.com/gocolly/colly/v2"
)

func scrape(search string) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only root url and urls which start with "e" or "h" on httpbin.org
	)
	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// Start scraping on http://httpbin.org
	c.Visit(search)
}

func main() {
	A := os.Args[1:]
	searchtext := strings.Join(A," ")
	helper := " 'index of/'" + " -unknownsecret.info" + " -wikipedia" + " -sirens.rocks" + " -billboard.com" + " -books.google.com"
	format := " .mp3"
	query := searchtext + helper + format
	r , err:=	googlesearch.Search(nil, query)
    
	if err != nil {
		panic(err)
	 }
	 for i := 0; i < len(r); i++ {
		if strings.Contains(r[i].Description, "mp3"){
			fmt.Println("Printing links from " + r[i].URL)
			// scrape(r[i].URL)
		}
	}

}

