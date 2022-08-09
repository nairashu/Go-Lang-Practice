package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	func() {
		fmt.Println("Ashish")
	}()

	fmt.Println("Working with Concurrency")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("Ashish")
	}()

	waitGroup.Wait()
}
