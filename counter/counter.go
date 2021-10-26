package counter

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sharedValue int
	lock        chan int
}

func NewCounter() *Counter {
	return &Counter{
		sharedValue: 0,
		lock:        make(chan int, 1),
	}
}

func (counter *Counter) Inc(id string, wg *sync.WaitGroup) int {
	counter.lock <- 0
	fmt.Println(id + " locked Counter under Inc ")
	counter.sharedValue++
	temp := counter.sharedValue
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(id + " will unlock Counter under Inc ")
	<-counter.lock
	defer wg.Done()
	return temp
}

func (counter *Counter) Dec(id string, wg *sync.WaitGroup) int {
	counter.lock <- 0
	fmt.Println(id + " locked Counter under Dec ")
	counter.sharedValue--
	temp := counter.sharedValue
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(id + " will unlock Counter under Dec ")
	<-counter.lock
	defer wg.Done()
	return temp
}