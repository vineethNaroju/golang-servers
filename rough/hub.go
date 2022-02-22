package main

import (
	"fmt"
	"sync"
)

func MainG() {

	recv := make(<-chan string)

	send := make(chan<- string)

	fmt.Printf("%T\n", recv)
	fmt.Printf("%T\n", send)

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(val int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("val:", val)
		}(i, wg)
	}

	wg.Wait()

	fmt.Println("all done")
}
