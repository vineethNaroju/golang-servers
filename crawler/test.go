package crawler

import (
	"fmt"
	"time"
)

func Test() {

	mp := make(map[string]int)

	mp["google"] = 1
	mp["meta"] = 2
	mp["netflix"] = 1
	mp["gmail"] = 2

	c := NewCrawler(mp)

	urls := []string {"google", "meta", "netflix", "gmail"}

	start := time.Now()

	c.crawlUrlsLinearly(urls)

	elapsed := time.Since(start)

	fmt.Println(elapsed)

	a := time.Now()

	c.crawlUrlsConcurrently(urls)

	b := time.Since(a)

	fmt.Println(b)
}