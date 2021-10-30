package crawler

import (
	"time"
)

type Crawler struct {
	urls map[string]int
}

func NewCrawler(urls map[string]int) *Crawler {
	return &Crawler{
		urls: urls,
	}
}

func (crawler *Crawler) crawlUrlsConcurrently(urls []string) {

	doneChannel := make(chan bool)

	for _, url := range urls {
		go func(c *Crawler, url string) {
			time.Sleep(1000 * time.Millisecond * time.Duration(crawler.urls[url]))
			doneChannel <- true
		}(crawler, url)
	}

	for i := 0; i < len(urls); i++ {
		<- doneChannel
	}
}

func (crawler *Crawler) crawlUrlsLinearly(urls []string) {
	for _, url := range urls {
		time.Sleep(1000 * time.Millisecond * time.Duration(crawler.urls[url]))
	}
}