package datasource

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type Stock struct {
	Name       string
	Percentage string
}

type Fund struct {
	CurrentPrice float64
	SourceURL    string
	Name         string
	Code         string
	StockList    []Stock
}

var host = "https://fund.eastmoney.com"

func GetFundsData(codes []string) []*Fund {
	result := []*Fund{}

	var urls []string
	for _, code := range codes {
		urls = append(urls, fmt.Sprintf("%s/%s.html", host, code))
	}

	ch := make(chan *Fund, len(urls))
	defer close(ch)

	c := createExtractCollector(ch)

	for _, url := range urls {
		c.Visit(url)
		c.Wait()
	}

	for {
		x := <-ch
		result = append(result, x)
		if len(result) == len(urls) {
			break
		}
	}

	return result
}

func GetFundsDataWithQueue(codes []string, queueLength int) []*Fund {
	result := []*Fund{}

	var urls []string
	for _, code := range codes {
		urls = append(urls, fmt.Sprintf("%s/%s.html", host, code))
	}

	ch := make(chan *Fund, len(urls))
	defer close(ch)

	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		queueLength, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c := createExtractCollector(ch)

	for i, url := range urls {
		// Add URLs to the queue
		q.AddURL(fmt.Sprintf("%s?n=%d", url, i))
	}
	// Consume URLs
	q.Run(c)

	for {
		x := <-ch
		result = append(result, x)
		if len(result) == len(urls) {
			break
		}
	}

	return result
}

func createExtractCollector(ch chan *Fund) *colly.Collector {
	c := colly.NewCollector()

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  fmt.Sprintf("*%s.*", host),
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnHTML("#body", func(e *colly.HTMLElement) {
		priceStr := e.DOM.Find("dl.dataItem02").ChildrenFiltered("dd.dataNums").Children().Eq(0).Text()
		price, _ := strconv.ParseFloat(priceStr, 64)
		// fmt.Println(price)

		info := e.DOM.Find("div.fundDetail-tit").Text()
		reg := regexp.MustCompile(`\((.*?)\)`)
		nameSplice := reg.Split(info, -1)
		name := nameSplice[0]
		codeSplice := reg.FindAllString(info, -1)
		code := strings.Trim(codeSplice[len(codeSplice)-1], "(")
		code = strings.Trim(code, ")")
		// fmt.Printf("%v, %v\n", nameSplice[0], code)

		stockQuery := "#quotationItem_DataTable > div.bd > ul > li.position_shares > .poptableWrap > table > tbody > tr"
		stocks := make([]Stock, 0, 10)
		e.DOM.Find(stockQuery).Each(func(index int, l *goquery.Selection) {
			if index == 0 {
				return
			}
			company := l.ChildrenFiltered("td:nth-child(1)").Text()
			percentage := l.ChildrenFiltered("td:nth-child(2)").Text()
			stocks = append(stocks, Stock{
				Name:       company,
				Percentage: percentage,
			})
			// fmt.Printf("%v, %v\n", company, percentage)
		})
		ch <- &Fund{
			CurrentPrice: price,
			Name:         name,
			Code:         code,
			StockList:    stocks,
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, ", Error:", err)
		ch <- nil
	})
	return c
}
