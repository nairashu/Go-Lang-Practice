package puzzle2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var scoreByStrategy := make(map[string]int)

func RunPuzzle2() {
	// open input file
	f, err := os.Open("./puzzle2/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanLine := bufio.NewScanner(f)

	
	totalScore := 0
	
	// Parse the file to scan each line and get the score for each row
	for scanLine.Scan() {
		strategy := scanLine.Text()
		score, ok := scoreByStrategy[strategy]
		if ok {
			totalScore = totalScore + score
		} else {
			calculateScoreForStrategy(strategy)
		}
	}
}

func calculateScoreForStrategy(string strategy) int {
	str := strings.Split(strategy, " ")
	log.Println("Opponent Played: ", str[0], " You should play: ", str[1])
	
	// Draw condition
	if str[0] == str[1] { 
		scoreByStrategy[strategy] = 3
		return 3
	}

	// Loosing conditions
	if str[0] == "A" && str[1] == "Z" {
		scoreByStrategy[strategy] = 3
		return 3
	}

	if str[0] == "B" && str[1] == "X" {
		scoreByStrategy[strategy] = 1
		return 1
	}

	if str[0] == "C" && str[1] == "Y" {
		scoreByStrategy[strategy] = 2
		return 2
	}

	// Winning conditions
	if str[0] == "A" && str[1] == "Y" {
		scoreByStrategy[strategy] = 8
		return 8
	}

	if str[0] == "B" && str[1] == "Z" {
		scoreByStrategy[strategy] = 9
		return 9
	}

	if str[0] == "C" && str[1] == "X" {
		scoreByStrategy[strategy] = 7
		return 7
	}
	
	log.Println("Invalid strategy")
	return
}