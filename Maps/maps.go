package main

import (
	"fmt"
)

func main() {
	fmt.Println("Working with Maps")

	leagueTitles := make(map[string]int)
	leagueTitles["ManUtd"] = 20
	leagueTitles["Liverpool"] = 19
	leagueTitles["Chelsea"] = 6
	leagueTitles["ManCity"] = 7

	fmt.Println("Length of leagueTitles map is ", len(leagueTitles))

	fmt.Printf("LeagueTitles map is %v \n", leagueTitles)

	for k, v := range leagueTitles {
		fmt.Println("Key: ", k, " and Value: ", v)
	}

	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
	}

	for k, v := range testMap {
		fmt.Printf("Key: %v and Value: %v \n", k, v)
	}

	fmt.Println("A = ", testMap["A"])
	testMap["A"] = 100
	fmt.Println("A = ", testMap["A"])
	fmt.Println(testMap)
	testMap["H"] = 8
	fmt.Println(testMap)
	delete(testMap, "C")
	fmt.Println(testMap)
}
