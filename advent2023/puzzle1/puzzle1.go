package puzzle1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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
