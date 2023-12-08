package puzzle6

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type RaceRecord struct {
	time     int
	distance int
}

func RunPuzzle6Solution1() {
	// dat, err := os.ReadFile("./puzzle6/Example1.txt")
	// dat, err := os.ReadFile("./puzzle6/Input1.txt")
	// dat, err := os.ReadFile("./puzzle6/Example2.txt")
	dat, err := os.ReadFile("./puzzle6/Input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	// Read race records
	raceRecords := readRaceRecords(data)

	// fmt.Printf("Part 1: %d\n", getMultiplicationOfValidOptionsToBreakRecord(raceRecords))

	fmt.Printf("Part 2: %d\n", getValidOptionCountIgnoringSpaces(raceRecords))
}

func getValidOptionCountIgnoringSpaces(raceRecords []RaceRecord) int {
	if len(raceRecords) != 1 {
		log.Fatal("Invalid data")
	}
	lowestTime := getTimeToPressButtonWithPointerAndIncrement(raceRecords[0], 0, 1)
	fmt.Println("Lowest Time: ", lowestTime)
	highestTime := getTimeToPressButtonWithPointerAndIncrement(raceRecords[0], raceRecords[0].time, -1)
	fmt.Println("Highest Time: ", highestTime)
	return highestTime - lowestTime + 1
}

func getTimeToPressButtonWithPointerAndIncrement(raceRecord RaceRecord, startpointer, incrementValue int) int {
	time := raceRecord.time
	distance := raceRecord.distance
	for startpointer <= time && startpointer >= 0 {
		timeRemaining := time - startpointer
		if timeRemaining*startpointer > distance {
			break
		}
		startpointer += incrementValue
	}
	return startpointer
}

func getMultiplicationOfValidOptionsToBreakRecord(raceRecords []RaceRecord) int {
	product := 1
	for _, raceRecord := range raceRecords {
		fmt.Printf("%d %d\n", raceRecord.time, raceRecord.distance)
		optionCount := findValidNumberOfOptionToBreakRecord(raceRecord)
		product *= optionCount
	}
	return product
}

func findValidNumberOfOptionToBreakRecord(raceRecord RaceRecord) int {
	validTimes := make([]int, 0)
	time := raceRecord.time
	distance := raceRecord.distance
	lowPtrSpeed := 0
	highPtrSpeed := time
	for lowPtrSpeed <= highPtrSpeed {
		lTimeRemaining := time - lowPtrSpeed
		hTimeRemaining := time - highPtrSpeed
		if lTimeRemaining*lowPtrSpeed > distance {
			validTimes = append(validTimes, lowPtrSpeed)
		}
		if hTimeRemaining*highPtrSpeed > distance && lowPtrSpeed != highPtrSpeed {
			validTimes = append(validTimes, highPtrSpeed)
		}
		lowPtrSpeed++
		highPtrSpeed--
	}

	sort.Ints(validTimes)
	fmt.Println("Valid Time Options to press the button: ", validTimes)
	return len(validTimes)
}

func readRaceRecords(data []string) []RaceRecord {
	re := regexp.MustCompile(`\d+`)
	times := re.FindAllString(data[0], -1)
	distances := re.FindAllString(data[1], -1)
	if len(times) != len(distances) {
		log.Fatal("Invalid data")
	}

	raceRecords := make([]RaceRecord, len(times))

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		raceRecord := RaceRecord{time: time, distance: distance}
		raceRecords[i] = raceRecord
	}
	return raceRecords
}
