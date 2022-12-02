package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func RunBufferedChannel() {

	fmt.Println("Working with Concurrency")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	fmt.Println("Working with Buffered channel")
	buffChan := make(chan string, 5)
	buffChan <- "Ashish"

	go func() {
		defer waitGroup.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
		buffChan <- "Tyler"
	}()

	go func() {
		defer waitGroup.Done()
		name := <-buffChan
		fmt.Println(name)
	}()

	waitGroup.Wait()
	pendingName := <-buffChan
	fmt.Println("Name left in the Buffered channel ia ", pendingName)
}
