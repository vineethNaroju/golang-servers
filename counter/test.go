package counter

import (
	"fmt"
	"sync"
)

func Test() {

	counter := NewCounter()

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