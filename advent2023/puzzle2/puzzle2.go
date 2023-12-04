package puzzle2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RunProblem1() {

	re := regexp.MustCompile(`\d+|blue|green|red`)
	// open file
	f, err := os.Open("./puzzle2/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	sum := 0
	redMax := 12
	greenMax := 13
	blueMax := 14
	// i := 0
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		fmt.Printf("line: %s\n", line)
		gameDetails := re.FindAllString(line, -1)
		fmt.Println("gameDetails:", gameDetails)
		gameIndex, err := strconv.Atoi(gameDetails[0])
		overboard := false
		if err != nil {
			log.Fatal(err)
		}

		for i := 1; i < len(gameDetails); i = i + 2 {
			colorValue, err := strconv.Atoi(gameDetails[i])
			if err != nil {
				log.Fatal(err)
			}
			color := gameDetails[i+1]
			if color == "red" && colorValue > redMax {
				fmt.Println("red value overboard")
				overboard = true
				break
			} else if color == "green" && colorValue > greenMax {
				fmt.Println("green value overboard")
				overboard = true
				break
			} else if color == "blue" && colorValue > blueMax {
				fmt.Println("blue value overboard")
				overboard = true
				break
			}
		}

		if !overboard {
			sum = sum + gameIndex
		}
	}

	fmt.Println("The sum of all of the calibration values is: ", sum)
}

func RunPuzzle2Solution2() {
	dat, err := os.ReadFile("./puzzle2/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	fmt.Printf("Part 1: %d\n", getValidGameIDSum(data))

	fmt.Printf("Part 2: %d\n", getValidGamePowerSum(data))
}

func getValidGamePowerSum(data []string) int {
	sum := 0
	for _, value := range data {
		gameDetails := extractGameDetails(value)
		minRed, minBlue, minGreen := getMinValuesForValidGame(gameDetails)
		power := minRed * minBlue * minGreen
		sum = sum + power
	}
	return sum
}

func getMinValuesForValidGame(gameDetails []string) (int, int, int) {
	minRed, minBlue, minGreen := 0, 0, 0
	for i := 1; i < len(gameDetails); i = i + 2 {
		colorValue, err := strconv.Atoi(gameDetails[i])
		if err != nil {
			log.Fatal(err)
		}
		color := gameDetails[i+1]
		if color == "red" && colorValue > minRed {
			minRed = colorValue
		} else if color == "green" && colorValue > minGreen {
			minGreen = colorValue
		} else if color == "blue" && colorValue > minBlue {
			minBlue = colorValue
		}
	}
	return minRed, minBlue, minGreen
}

func getValidGameIDSum(data []string) int {
	sum := 0
	for _, value := range data {
		gameDetails := extractGameDetails(value)
		gameIndex, err := strconv.Atoi(gameDetails[0])
		if err != nil {
			log.Fatal(err)
		}

		if isValidGame(gameDetails, 12, 14, 13) {
			sum = sum + gameIndex
		}
	}
	return sum
}

func isValidGame(gameDetails []string, redMax, blueMax, greenMax int) bool {
	for i := 1; i < len(gameDetails); i = i + 2 {
		colorValue, err := strconv.Atoi(gameDetails[i])
		if err != nil {
			log.Fatal(err)
		}
		color := gameDetails[i+1]
		if color == "red" && colorValue > redMax {
			fmt.Println("red value overboard")
			return false
		} else if color == "green" && colorValue > greenMax {
			fmt.Println("green value overboard")
			return false
		} else if color == "blue" && colorValue > blueMax {
			fmt.Println("blue value overboard")
			return false
		}
	}
	return true
}

func extractGameDetails(value string) []string {
	re := regexp.MustCompile(`\d+|blue|green|red`)
	gameDetails := re.FindAllString(value, -1)
	return gameDetails
}

func RunPuzzle2() {
}
