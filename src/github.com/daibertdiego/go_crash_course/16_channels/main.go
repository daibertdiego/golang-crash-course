package main

import (
	"fmt"
	"time"
)

func main() {
	intChannel := make(chan int)
	go giveMeNumber(intChannel)
	doubleNumber(intChannel)
}

func giveMeNumber(intChannel chan int) {
	defer close(intChannel) // defer -> executes when it finishes
	for i := 0; i <= 10; i++ {
		intChannel <- i
	}
}

func doubleNumber(intChannel chan int) {
	// for number := range intChannel {
	for {
		number, isOpen := <-intChannel
		if !isOpen {
			break
		}
		fmt.Println(number, "Double is:", number*2)
		time.Sleep(time.Millisecond * 500)
	}
}
