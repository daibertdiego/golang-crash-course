package main

import (
	"fmt"
	"time"
)

func doSomething(thingToDo string) {
	for i := 0; i <= 10; i++ {
		fmt.Println("Step: ", i, "I am ", thingToDo)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go doSomething("eating")           // Async
	go doSomething("watching netflix") // Async
	go doSomething("drinking")

	// Wait for a user input and force the main thread to wait
	// while the other threds are running
	fmt.Scanln()

}
