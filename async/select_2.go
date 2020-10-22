package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch2 := make(chan int, 2)
	ch2 <- 3
	// LOOP:
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("chan1 val", v1)
		case v2 := <-ch2:
			fmt.Println("chan2 val", v2)
		default:
			fmt.Println("break LOOP")
			break LOOP
			// break (this only break select, not for!)
		}
	}
}
