package main

import (
	"fmt"
)

func main() {
	cancelCh := make(chan struct{}) // empty struct is just a sign, means something.
	dataCh := make(chan int)

	go func(cancelCh chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case <-cancelCh:
				return
			case dataCh <- val:
				fmt.Println("Here I push:", val)
				val++
				// default:
				// 	fmt.Println("default select")
			}
		}
	}(cancelCh, dataCh)

	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- struct{}{}
			break
		} else {
			fmt.Println("curVal: ", curVal)
		}
	}

}
