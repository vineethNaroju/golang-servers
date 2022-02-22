package main

import (
	"fmt"
	"time"
)

func portalX(cx chan string) {
	for i := 0; i <= 3; i++ {
		cx <- fmt.Sprintf("x:%d", i)
	}
}

func portalY(cy chan string) {
	cy <- "y"
}

func MainM() {

	cx := make(chan string)
	cy := make(chan string)

	go portalX(cx)
	go portalY(cy)

	var tempChannel chan int // = make(chan int)

	fmt.Println(tempChannel)
	fmt.Printf("%T\n", tempChannel)

	go func() {
		defer fmt.Println("hahah")
	}()

	time.Sleep(10 * time.Second)

	fmt.Println("bye bye")

	for {
		select {
		case x := <-cx:
			fmt.Println(x)
		case y := <-cy:
			fmt.Println(y)
		default:
			fmt.Println("Nothing")
		}
	}
}
