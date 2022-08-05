package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "ashish nair"
	course := "kubernetes with go lang"

	fmt.Println(convertor(name, course))

	topScore := getTopScore(67, 98, 21, 22, 5, 99, 67)
	fmt.Println("Top Score is ", topScore)
}

func convertor(s1, s2 string) (str1, str2 string) {
	s1 = strings.ToUpper(s1)
	s2 = strings.Title(s2)

	return s1, s2
}

func getTopScore(scores ...int) int {
	topScore := scores[0]
	for _, i := range scores {
		if i > topScore {
			topScore = i
		}
	}
	return topScore
}
