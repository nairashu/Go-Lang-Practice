package puzzle1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func RunPuzzle1() {
	// open file
	f, err := os.Open("./puzzle1/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	caloriesByElves := make(map[int]int)
	calories := make([]int, 0, 10)
	calorieTotal := 0
	elfNumber := 1
	maxCalorie := 0

	for scanner.Scan() {
		// do something with a line
		//fmt.Printf("line: %s\n", scanner.Text())
		if scanner.Text() == "" {
			caloriesByElves[elfNumber] = calorieTotal
			calories = append(calories, calorieTotal)
			if maxCalorie < calorieTotal {
				maxCalorie = calorieTotal
			}
			calorieTotal = 0
			elfNumber++
		} else {
			curr, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			calorieTotal = calorieTotal + curr
		}
	}

	fmt.Println("Max Calorie is: ", maxCalorie)

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	top3Total := 0
	for i := 0; i < 3; i++ {
		top3Total = top3Total + calories[i]
	}

	fmt.Println("Total of top 3 Max Calorie is: ", top3Total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
