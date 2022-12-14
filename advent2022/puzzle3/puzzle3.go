package puzzle3

import (
	"bufio"
	"errors"
	"log"
	"os"
)

// var priorityorder = map[rune]int{
// 	'a': 1,
// 	'b': 2,
// 	'c': 3,
// 	'd': 4,
// 	'e': 5,
// 	'f': 6,
// 	'g': 7,
// 	'h': 8,
// 	'i': 9,
// 	'j': 10,
// 	'k': 11,
// 	'l': 12,
// 	'm': 13,
// 	'n': 14,
// 	'o': 15,
// 	'p': 16,
// 	'q': 17,
// 	'r': 18,
// 	's': 19,
// 	't': 20,
// 	'u': 21,
// 	'v': 22,
// 	'w': 23,
// 	'x': 24,
// 	'y': 25,
// 	'z': 26,
// 	'A': 27,
// 	'B': 28,
// 	'C': 29,
// 	'D': 30,
// 	'E': 31,
// 	'F': 32,
// 	'G': 33,
// 	'H': 34,
// 	'I': 35,
// 	'J': 36,
// 	'L': 37,
// 	'M': 38,
// 	'N': 39,
// 	'O': 40,
// 	'P': 41,
// 	'Q': 42,
// 	'R': 43,
// 	'S': 44,
// 	'T': 45,
// 	'U': 46,
// 	'V': 47,
// 	'W': 48,
// 	'X': 49,
// 	'Y': 50,
// 	'Z': 51,
// }

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
		//log.Println(rackSackInputLine)
		rackSackChars := []rune(rackSackInputLine)

		misplacedItems := identifyMisplacedItemTypes(rackSackChars)
		log.Printf("%c ", misplacedItems)
		rackSackPriority := calculatePriorityTotalPerRackSack(misplacedItems)
		totalPriority = totalPriority + rackSackPriority
		//log.Println("Total Priority is: ", totalPriority)
	}

	log.Println("Total Priority is: ", totalPriority)
}

func RunPuzzle3Part2() {

	// Read the Input File
	f, err := os.Open("./puzzle3/Input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Remember to close the file at the end of the program
	defer f.Close()

	// Assign the scanner to read the file line by line
	scanLine := bufio.NewScanner(f)
	var ruckSackInputLines []string

	for scanLine.Scan() {
		rackSackInputLine := scanLine.Text()
		ruckSackInputLines = append(ruckSackInputLines, rackSackInputLine)
	}

	totalSum := 0
	for i := 0; i < len(ruckSackInputLines); i = i + 3 {
		badgeItem, err := identifyCommonChar(ruckSackInputLines[i], ruckSackInputLines[i+1], ruckSackInputLines[i+2])
		if err != nil {
			log.Fatal(err)
		} else {
			badgeItemPriority := calculatePriorityTotalPerRackSack([]rune{badgeItem})
			totalSum = totalSum + badgeItemPriority
		}
	}

	log.Println("Total Sum is: ", totalSum)
}

func identifyCommonChar(s1, s2, s3 string) (rune, error) {
	log.Println(s1, s2, s3)
	str1 := []rune(s1)
	str2 := []rune(s2)
	str3 := []rune(s3)

	// len1 := sortRuneString(str1).Len()
	// len2 := sortRuneString(str2).Len()
	// len3 := sortRuneString(str3).Len()
	// log.Println(len1, len2, len3)

	// maxLen := math.Max(float64(len1), float64(len2))
	// maxLen = math.Max(maxLen, float64(len3))

	for i := 0; i < len(str1); i++ {
		if SliceContains(str2, str1[i]) && SliceContains(str3, str1[i]) {
			return str1[i], nil
		}
	}

	return rune('$'), errors.New("No common char found in ruckSack")
}

// type sortRuneString []rune

// func (s sortRuneString) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

// func (s sortRuneString) Less(i, j int) bool {
// 	return s[i] < s[j]
// }

// func (s sortRuneString) Len() int {
// 	return len(s)
// }

func calculatePriorityTotalPerRackSack(misplacedItems []rune) int {
	total := 0
	checkedItem := make(map[rune]bool)

	for _, v := range misplacedItems {
		if _, ok := checkedItem[v]; !ok {
			checkedItem[v] = true
			total = total + int(priority(v))
		}
	}

	log.Println("Racksack Total Priority: ", total)
	return total
}

func identifyMisplacedItemTypes(rackSackChars []rune) []rune {
	intersection := make([]rune, 0)
	size := len(rackSackChars)

	comp1 := rackSackChars[0 : size/2]
	comp2 := rackSackChars[size/2 : size]
	//log.Printf("%c ", comp1)
	//log.Printf("%c ", comp2)

	m := make(map[rune]bool)

	for _, v := range comp1 {
		m[v] = true
	}

	for _, v := range comp2 {
		if _, ok := m[v]; ok {
			if !SliceContains(intersection, v) {
				intersection = append(intersection, v)
				return intersection
			}
		}
	}
	return intersection
}

func SliceContains(intersection []rune, v rune) bool {
	for _, val := range intersection {
		if val == v {
			return true
		}
	}
	return false
}

func priority(item rune) uint64 {
	if item >= 'a' && item <= 'z' {
		return uint64(item-'a') + 1
	}
	return uint64(item-'A') + 27
}

/*******

Solution from reddit channels
*****/

/*
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunPuzzle3() {
	var sum uint64
	// Read the Input File
	f, err := os.Open("./puzzle3/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Remember to close the file at the end of the program
	defer f.Close()

	// Assign the scanner to read the file line by line
	scanLine := bufio.NewScanner(f)

	for scanLine.Scan() {
		rackSackInputLine := scanLine.Text()
		var set uint64
		for i, item := range rackSackInputLine {
			prio := priority(item)
			mask := uint64(1) << prio
			if i < len(rackSackInputLine)/2 {
				set |= mask
			} else if set&mask == mask {
				sum += prio
				set &^= mask // delete
			}
		}
	}
	fmt.Println(sum)
}

func priority(item rune) uint64 {
	if item >= 'a' && item <= 'z' {
		return uint64(item-'a') + 1
	}
	return uint64(item-'A') + 27
}
*/
