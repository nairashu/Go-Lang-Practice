package puzzle3

import (
	"bufio"
	"log"
	"os"
)

var priorityorder = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

// var pr = []rune { 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l','m','n','o','p','q','r','s','t','u','v','w','x','y','z','A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'}

func RunPuzzle3() {

	// Read the Input File
	f, err := os.Open("./puzzle3/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Remember to close the file at the end of the program
	defer f.Close()

	// Assign the scanner to read the file line by line
	scanLine := bufio.NewScanner(f)
	totalPriority := 0

	for scanLine.Scan() {
		rackSackInputLine := scanLine.Text()
		rackSackChars := []rune(rackSackInputLine)

		misplacedItems := identifyMisplacedItemTypes(rackSackChars)
		rackSackPriority := calculatePriorityTotalPerRackSack(misplacedItems)
		totalPriority = totalPriority + rackSackPriority
	}

	log.Println("Total Priority is: ", totalPriority)
}

func calculatePriorityTotalPerRackSack(misplacedItems []rune) int {
	total := 0
	checkedItem := make(map[rune]bool)

	for _, v := range misplacedItems {
		if _, ok := checkedItem[v]; !ok {
			checkedItem[v] = true
			total = total + priorityorder[v]
		}
	}

	log.Println("Racksack Total Priority: ", total)
	return total
}

func identifyMisplacedItemTypes(rackSackChars []rune) (intersection []rune) {
	size := len(rackSackChars)

	comp1 := rackSackChars[0 : size/2]
	comp2 := rackSackChars[size/2 : size]
	log.Printf("%c ", comp1)
	log.Printf("%c ", comp2)

	m := make(map[rune]bool)

	for _, v := range comp1 {
		m[v] = true
	}

	for _, v := range comp2 {
		if _, ok := m[v]; ok {
			intersection = append(intersection, v)
		}
	}
	return
}
