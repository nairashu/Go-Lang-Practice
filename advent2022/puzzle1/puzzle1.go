package puzzle1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunPuzzle1() {
	// open file
	f, err := os.Open("Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
