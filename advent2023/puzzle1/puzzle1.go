package puzzle1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var digitConversion = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "zero": 0}

func RunPuzzle1() {
	re := regexp.MustCompile(`\d`)

	// open file
	f, err := os.Open("./puzzle1/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		fmt.Printf("line: %s\n", line)
		numbers := re.FindAllString(line, -1)
		fmt.Println("numbers:", numbers)
		size := len(numbers)
		// No Numbers found on the line so nothing to add to the sum
		if size == 0 {
			fmt.Println("No numbers found in the line")
			continue
		}
		var lineCallibration string
		// Single number found on the line so the number is concatenate the digit to itself and add to the sum
		if size == 1 {
			fmt.Println("Single number found in the line")
			lineCallibration = numbers[0] + numbers[0]
		}
		lineCallibration = numbers[0] + numbers[size-1]
		fmt.Println("lineCallibration:", lineCallibration)
		lineValue, err := strconv.Atoi(lineCallibration)
		if err != nil {
			log.Fatal(err)
		}
		sum += lineValue
	}

	fmt.Println("The sum of all of the calibration values is: ", sum)
}

func RunPuzzle2() {
	re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine|zero`)

	// open file
	//f, err := os.Open("./puzzle1/Example2.txt")
	f, err := os.Open("./puzzle1/Input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		// fmt.Printf("line: %s\n", line)
		numbers := re.FindAllString(line, -1)
		//fmt.Println("numbers:", numbers)
		size := len(numbers)
		// No Numbers found on the line so nothing to add to the sum
		if size == 0 {
			// fmt.Println("No numbers found in the line")
			continue
		}
		// Single number found on the line so the number is concatenate the digit to itself and add to the sum
		if size == 1 {
			// fmt.Println("Single number found in the line")
			lineValue := findLineValue(numbers[0], numbers[0])
			sum += lineValue
			fmt.Println("sum:", sum)
			continue
		}
		multiLineValue := findLineValue(numbers[0], numbers[size-1])
		sum += multiLineValue
		fmt.Println("sum:", sum)
	}

	fmt.Println("The sum of all of the calibration values is: ", sum)
}

func findLineValue(s1, s2 string) int {
	digit1 := s1
	digit2 := s2
	if len(s1) != 1 {
		digit1 = strconv.Itoa(digitConversion[s1])
	}

	if len(s2) != 1 {
		digit2 = strconv.Itoa(digitConversion[s2])
	}

	lineCallibration := digit1 + digit2
	fmt.Println("lineCallibration:", lineCallibration)
	lineValue, err := strconv.Atoi(lineCallibration)
	if err != nil {
		log.Fatal(err)
	}
	return lineValue
}
