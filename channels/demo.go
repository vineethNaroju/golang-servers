package channels

import (
	"fmt"
	"time"
)

func Demo() {
	chan1 := make(chan string)
	chan2 := make(chan int)
	done:= 0

	go func() {
		for i:=0 ; i<5; i++ {
			chan1 <- fmt.Sprintf("hello x %d", i)
		}
		done++
	}()

	go func() {
		for i:=0; i<5; i++ {
			chan2 <- (i*100)
		}
		done++
	}()

	for done < 2 {
		select {
			case s := <- chan1 :
				fmt.Println(s)
			case n := <- chan2 :
				fmt.Println(n)
			default:
		}
	}



	time.Sleep(1000 * time.Millisecond)
}
