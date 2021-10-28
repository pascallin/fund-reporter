package datasource

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

var urls = []string{
	"https://fund.eastmoney.com/161725.html",
}

type Fund struct {
	CurrentPrice float64
	SourceURL    string
}

func GetFundsCurrentPrice() {
	for _, url := range urls {
		c := colly.NewCollector()
		c.OnHTML("dl.dataItem02", func(e *colly.HTMLElement) {
			price := e.DOM.ChildrenFiltered("dd.dataNums").Children().Eq(0).Text()
			fmt.Println(price)
		})
		c.Visit(url)
	}
}

func GetFundsCurrentPriceSerial() {
	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c := colly.NewCollector()
	c.OnHTML("dl.dataItem02", func(e *colly.HTMLElement) {
		price := e.DOM.ChildrenFiltered("dd.dataNums").Children().Eq(0).Text()
		fmt.Println(price)
	})

	for i, url := range urls {
		// Add URLs to the queue
		q.AddURL(fmt.Sprintf("%s?n=%d", url, i))
	}
	// Consume URLs
	q.Run(c)
}
