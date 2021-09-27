package main

import (
    "fmt"
    "os"
	"strings"

	"github.com/rocketlaunchr/google-search"
	"github.com/gocolly/colly/v2"
)

func findMusic() {
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

