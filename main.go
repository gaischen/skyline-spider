package main

import (
	"fmt"
	"github.com/skyline/skyline-spider/agent"
	"github.com/skyline/skyline-spider/poster"
)
import colly "github.com/gocolly/colly/v2"

func main() {
	fmt.Println("test")
	c := colly.NewCollector(
		colly.AllowedDomains("www.shein.com"),
		colly.UserAgent(agent.GetChromeUserAgent()),
		colly.AllowURLRevisit(),
	)
	// On every a element which has href attribute call callback
	c.OnHTML("span", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping
	//err := c.Visit("https://www.shein.com/BLUES-High-Waisted-Jeggings-Without-Pocket-p-1587411-cat-1934.html?scici=navbar_WomenHomePage~~tab01navbar05~~5~~webLink~~SPcCccWomenCategory~~0~~50001")
	//if err != nil {
	//	fmt.Println(err)
	//}

	p := poster.NewHttpPoster()

	p.OnHtml("href", func(content string) {
		fmt.Println(content)
	})
	p.Call("http://www.baidu.com", nil)

}
