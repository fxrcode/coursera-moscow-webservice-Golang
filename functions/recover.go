package main

import (
	"fmt"
)

func deferTest1() {
	fmt.Println("Some useful work")
	panic("something bad happened")
	return
}

func deferTest2() {
	defer func() {
		fmt.Println("in panic mode")
	}()
	fmt.Println("Some useful work")
	panic("something bad happened")
	return
}

func deferTest3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend:", err)
		}
	}()
	fmt.Println("Some useful work")
	panic("something bad happened")
	return
}

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend FIRST:", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend SECOND:", err)
			// panic(2)
			panic("second panic")
		}
	}()
	fmt.Println("Some userful work")
	panic("something bad happend")
	return
}

func main() {
	// deferTest1()
	// deferTest2()
	// deferTest3()
	deferTest()
	return
}
