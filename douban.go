package main

import (
	"os"
	"github.com/anaskhan96/soup"
	"fmt"
)

func main() {
	resp, err := soup.Get("https://movie.douban.com/subject/27605698/")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("a", "rel", "v:starring")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}

}
