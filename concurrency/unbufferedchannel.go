package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func RunUnbufferedChannel() {
	fmt.Println("Working with Concurrency")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	fmt.Println("Working with UnBuffered channel")
	unBuffChan := make(chan string)
	go func() {
		defer waitGroup.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
		unBuffChan <- "Tyler"
	}()

	go func() {
		defer waitGroup.Done()
		name := <-unBuffChan
		fmt.Println(name)
	}()

	waitGroup.Wait()
}
