package main

import (
	"strings"

	"github.com/rocketlaunchr/google-search"
	//"github.com/gocolly/colly/v2"
)

func findMusic(search string) string {
	searchtext := search + " "
	helper := " 'index of/'" + " -unknownsecret.info" + " -wikipedia" + " -sirens.rocks" + " -billboard.com" + " -books.google.com"
	format := " .mp3"
	query := searchtext + helper + format
	r , err:=	googlesearch.Search(nil, query)
    
	if err != nil {
		panic(err)
	 }

	 result := ""
	 for i := 0; i < len(r); i++ {
		if strings.Contains(r[i].Description, "mp3"){
			result = result + r[i].URL + "\n"
			// scrape(r[i].URL)
		}
	}
	return result
}