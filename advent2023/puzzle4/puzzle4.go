package puzzle4

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func RunPuzzle4Solution1() {
	// dat, err := os.ReadFile("./puzzle4/Example1.txt")
	dat, err := os.ReadFile("./puzzle4/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	fmt.Printf("Part 1: %d\n", getSumOfScratchcardPoints(data))

	fmt.Printf("Part 2: %d\n", calculateSumOfAllCopiesOfScratchcards(data))
}

var NoOfCardsByCard = make(map[int]int)
var WinningPointsByCard = make(map[int]int)

func calculateSumOfAllCopiesOfScratchcards(data []string) int {
	cardDeets := initialilzeCardDataAndExtractDetails(data)
	findWinningPointsForEachCard(cardDeets)
	// fmt.Println(WinningPointsByCard)
	updateCardCounts()
	// fmt.Println(NoOfCardsByCard)
	return calculateTotalCards()
}

func calculateTotalCards() int {
	totalCards := 0
	for _, count := range NoOfCardsByCard {
		totalCards = totalCards + count
	}
	return totalCards
}

func updateCardCounts() {
	for index := 0; index < len(NoOfCardsByCard); index++ {
		currentCardCount := NoOfCardsByCard[index]
		points := WinningPointsByCard[index]
		if points == 0 {
			continue
		}
		for j := 1; j <= points; j++ {
			NoOfCardsByCard[index+j] = NoOfCardsByCard[index+j] + currentCardCount
		}
		fmt.Println("Index: ", index, " Points: ", points, " CurrentCardCount: ", currentCardCount)
		fmt.Println(NoOfCardsByCard)
	}
}

func findWinningPointsForEachCard(cardDeets []string) {
	for i, cardInfo := range cardDeets {

		numbers := strings.Split(cardInfo, "|")
		re := regexp.MustCompile(`\d+`)
		winningNumbers := re.FindAllString(numbers[0], -1)
		yourNumbers := re.FindAllString(numbers[1], -1)
		sort.Strings(winningNumbers)
		// fmt.Println(winningNumbers)
		// fmt.Println(yourNumbers)

		matchCount := 0
		for _, num := range yourNumbers {
			i := sort.SearchStrings(winningNumbers, num)
			if i < len(winningNumbers) && winningNumbers[i] == num {
				// fmt.Println("You won!")
				matchCount++
			}
		}
		WinningPointsByCard[i] = matchCount
		// fmt.Println("Card: ", i, " Matches: ", matchCount)
	}
}

func initialilzeCardDataAndExtractDetails(data []string) []string {
	cardDeets := make([]string, len(data))
	for i, cardInfo := range data {
		info := strings.Split(cardInfo, ":")
		cardDeets[i] = info[1]
		NoOfCardsByCard[i] = 1
		WinningPointsByCard[i] = 0
	}
	return cardDeets
}

func getSumOfScratchcardPoints(data []string) int {
	totalPoints := 0

	for _, cardInfo := range data {
		fmt.Println(cardInfo)
		cardDets := strings.Split(cardInfo, ":")
		numbers := strings.Split(cardDets[1], "|")
		re := regexp.MustCompile(`\d+`)
		winningNumbers := re.FindAllString(numbers[0], -1)
		yourNumbers := re.FindAllString(numbers[1], -1)
		sort.Strings(winningNumbers)

		cardPoints := 0
		for _, num := range yourNumbers {
			i := sort.SearchStrings(winningNumbers, num)

			if i < len(winningNumbers) && winningNumbers[i] == num {
				fmt.Println("You won!")
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints = cardPoints * 2
				}
			}
		}
		totalPoints = totalPoints + cardPoints
	}
	return totalPoints
}
