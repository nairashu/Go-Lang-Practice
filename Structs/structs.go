package main

import (
	"fmt"
)

func main() {
	fmt.Println("Working with Structs")

	type footballClub struct {
		name           string
		playerCount    int
		staffCount     int
		avgEmployeeAge float64
	}

	var barcelona footballClub
	chelsea := new(footballClub)
	fmt.Println("Barcelona Details are ", barcelona)
	fmt.Println("Chelsea Details are ", chelsea)

	liverpool := footballClub{
		name:           "Liverpool F.C.",
		playerCount:    27,
		staffCount:     15,
		avgEmployeeAge: 32.2,
	}

	fmt.Printf("Liveropol Details are %v \n", liverpool)
	fmt.Printf("Number of Players in Liverpool are %v \n", liverpool.playerCount)
	liverpool.playerCount = 30
	fmt.Printf("Number of Players in Liverpool are %v \n", liverpool.playerCount)
}
