package main

import (
	"fmt"
	"sync"

	Counter "example.com/gobrr/counter"
)

func main() {
	counter := Counter.NewCounter()

	wg := &sync.WaitGroup{}
	
	go counter.Inc("a", wg)
	go counter.Inc("b", wg)
	go counter.Inc("c", wg)

	go counter.Dec("d", wg)
	go counter.Dec("e", wg)
	go counter.Dec("f", wg)

	wg.Add(6)
	wg.Wait()

	fmt.Println("All done")

	//time.Sleep(time.Millisecond * 10000)
}