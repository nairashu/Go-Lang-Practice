package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	switch "Kubernetes" {
	case "kubernetes":
		fmt.Println("Case 1 with lower k kubernetes")
	case "Kubernetes":
		fmt.Println("Case 2 with upper K Kubernetes")
		fallthrough
	case "Docker":
		fmt.Println("Case 3 with Docker")
	case "Terra":
		fmt.Println("Case 4 with Terra")
	default:
		fmt.Println("Invalid entry. Did not match any case")
	}

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("Number is even")
	case 1, 3, 5, 7, 9:
		fmt.Println("Number is odd")
	default:
		fmt.Println("Invalid entry. Did not match any case")
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
