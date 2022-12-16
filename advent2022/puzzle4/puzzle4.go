package puzzle4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunPuzzle4() {
	f, err := os.Open("./puzzle4/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	count := 0
	overlapCount := 0
	for scanner.Scan() {
		scanLine := scanner.Text()
		log.Println("Range Pair: ", scanLine)
		ranges := strings.Split(scanLine, ",")

		coll1 := strings.Split(ranges[0], "-")
		coll2 := strings.Split(ranges[1], "-")

		lcoll1, _ := strconv.Atoi(coll1[0])
		hcoll1, _ := strconv.Atoi(coll1[1])
		lcoll2, _ := strconv.Atoi(coll2[0])
		hcoll2, _ := strconv.Atoi(coll2[1])

		if checkCompleteSubset(lcoll1, lcoll2, hcoll1, hcoll2) {
			count++
		}

		if checkRangesOverlap(lcoll1, lcoll2, hcoll1, hcoll2) {
			overlapCount++
		}
	}
	log.Println("Count of complete subsets: ", count)

	log.Println("Count of Overlapping subsets: ", overlapCount)
}

func checkRangesOverlap(lcoll1, lcoll2, hcoll1, hcoll2 int) bool {
	if lcoll1 <= lcoll2 && lcoll2 <= hcoll1 {
		log.Println("Range 2 Overrides Range 1")
		return true
	} else if lcoll2 <= lcoll1 && lcoll1 <= hcoll2 {
		log.Println("Range 1 Overrides Range 2")
		return true
	} else {
		log.Println("Does not override")
		return false
	}
}

func checkCompleteSubset(lcoll1, lcoll2, hcoll1, hcoll2 int) bool {
	if lcoll1 == lcoll2 && hcoll1 == hcoll2 {
		log.Println("Same Range")
		return true
	} else if lcoll1 <= lcoll2 && hcoll1 >= hcoll2 {
		log.Println("Second Range is complete subset of first")
		return true
	} else if lcoll2 <= lcoll1 && hcoll2 >= hcoll1 {
		log.Println("First Range is complete subset of first")
		return true
	} else {
		log.Println("Not a Subset")
		return false
	}
}
