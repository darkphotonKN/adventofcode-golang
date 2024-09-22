package concurrency

import (
	"fmt"
	"time"
)

var ch = make(chan string)
var ch2 = make(chan string)

var canExit = make(chan bool)

func AdvancedConcurrency() {
	fmt.Println("Advanced Concurrency")
	fmt.Println("Waiting for my channel message")

	go writeLoop()

	// wg.Add(2)
	go readLoop()

	// wg.Wait()
	<-canExit
	fmt.Println("Waiting Done")
}

func writeLoop() {
	time.Sleep(time.Second * 1)

	ch <- "hey channel 1!!!!"

	time.Sleep(time.Second * 2)

	ch2 <- "for channel 2"

	time.Sleep(time.Second * 4)

	canExit <- true
}

func readLoop() {
	for {
		fmt.Println("waiting for messages to be sent to channels")
		select {
		case v := <-ch:
			fmt.Println("value received for CHANNEL ONE:", v)
			// close(ch)
			continue
		case v2 := <-ch2:
			fmt.Println("value received for CHANNEL TWO:", v2)
			// close(ch2)
			continue
		}
	}
}
