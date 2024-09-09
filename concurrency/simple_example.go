package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var messageChan = make(chan string)

func SimpleExample() {
	var wg sync.WaitGroup
	start := time.Now()

	wg.Add(3) // wait for 2 things

	go LoggerLong(&wg)
	go Logger2(&wg)
	go Logger(&wg)

	go listenToLoggers()

	wg.Wait() // wait for 2 done's before unlocking
	close(messageChan)

	end := time.Since(start)
	fmt.Printf("final time: %v\n", end)
}

func Logger(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	messageChan <- "Logger 1 is done"
}

func LoggerLong(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 4)
	messageChan <- "Logger 4 is done"
}

func Logger2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	messageChan <- "Logger 2 is done"
}

func listenToLoggers() {
	for message := range messageChan {
		fmt.Println(message)
	}
}
