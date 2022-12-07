package puzzle2

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
	scoreByStrategyLDW := make(map[string]int)
	scoreByStrategyRPS := make(map[string]int)

	totalScoreLDW := 0
	totalScoreRPS := 0

	for scanLine.Scan() {
		strategy := scanLine.Text()
		scoreLDW, ok := scoreByStrategyLDW[strategy]
		scoreRPS, ok := scoreByStrategyRPS[strategy]
		str := strings.Split(strategy, " ")
		log.Println("Opponent Played: ", str[0], " Your strategy play is: ", str[1])
		if ok {
			totalScoreLDW = totalScoreLDW + scoreLDW
			totalScoreRPS = totalScoreRPS + scoreRPS
			log.Println("ScoreLDW is: ", scoreLDW, " TotalLDW: ", totalScoreLDW, " ScoreRPS: ", scoreRPS, " TotalRPS: ", totalScoreRPS)
		} else {
			calculateScoreForStrategyWhereSecondColumnIsLDW(strategy, scoreByStrategyLDW)
			calculateScoreForStrategyWhereSecondColumnIsRPS(strategy, scoreByStrategyRPS)
			newScoreLDW := scoreByStrategyLDW[strategy]
			totalScoreLDW = totalScoreLDW + newScoreLDW
			newScoreRPS := scoreByStrategyRPS[strategy]
			totalScoreRPS = totalScoreRPS + newScoreRPS
			log.Println("NewScoreLDW is: ", newScoreLDW, " TotalLDW: ", totalScoreLDW, " NewScoreRPS: ", newScoreRPS, " TotalRPS: ", totalScoreRPS)
		}
	}
	log.Println("Total Strategy score for me when second column is Lose Draw Win: ", totalScoreLDW)
	log.Println("Total Strategy score for me when second column is Rock Paper Scissor: ", totalScoreRPS)

}

func calculateScoreForStrategyWhereSecondColumnIsRPS(strategy string, scoreByStrategy map[string]int) {
	str := strings.Split(strategy, " ")
	// log.Println("Opponent Played: ", str[0], " You should play: ", str[1])

	// Draw conditions
	if str[0] == "A" && str[1] == "X" {
		scoreByStrategy[strategy] = 4
	}

	if str[0] == "B" && str[1] == "Y" {
		scoreByStrategy[strategy] = 5
	}

	if str[0] == "C" && str[1] == "Z" {
		scoreByStrategy[strategy] = 6
	}

	// Loosing conditions
	if str[0] == "A" && str[1] == "Z" {
		scoreByStrategy[strategy] = 3
	}

	if str[0] == "B" && str[1] == "X" {
		scoreByStrategy[strategy] = 1
	}

	if str[0] == "C" && str[1] == "Y" {
		scoreByStrategy[strategy] = 2
	}

	// Winning conditions
	if str[0] == "A" && str[1] == "Y" {
		scoreByStrategy[strategy] = 8
	}

	if str[0] == "B" && str[1] == "Z" {
		scoreByStrategy[strategy] = 9
	}

	if str[0] == "C" && str[1] == "X" {
		scoreByStrategy[strategy] = 7
	}

	//log.Println("Score is: ", scoreByStrategy[strategy])
}

func calculateScoreForStrategyWhereSecondColumnIsLDW(strategy string, scoreByStrategy map[string]int) {
	str := strings.Split(strategy, " ")
	// log.Println("Opponent Played: ", str[0], " You should play: ", str[1])

	// Draw conditions
	if str[0] == "A" && str[1] == "Y" {
		scoreByStrategy[strategy] = 4
	}

	if str[0] == "B" && str[1] == "Y" {
		scoreByStrategy[strategy] = 5
	}

	if str[0] == "C" && str[1] == "Y" {
		scoreByStrategy[strategy] = 6
	}

	// Loosing conditions
	if str[0] == "A" && str[1] == "X" {
		scoreByStrategy[strategy] = 3
	}

	if str[0] == "B" && str[1] == "X" {
		scoreByStrategy[strategy] = 1
	}

	if str[0] == "C" && str[1] == "X" {
		scoreByStrategy[strategy] = 2
	}

	// Winning conditions
	if str[0] == "A" && str[1] == "Z" {
		scoreByStrategy[strategy] = 8
	}

	if str[0] == "B" && str[1] == "Z" {
		scoreByStrategy[strategy] = 9
	}

	if str[0] == "C" && str[1] == "Z" {
		scoreByStrategy[strategy] = 7
	}

	//log.Println("Score is: ", scoreByStrategy[strategy])
}
