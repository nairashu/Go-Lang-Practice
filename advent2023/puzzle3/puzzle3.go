package puzzle3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Position struct {
	x, y int
}

type Parts struct {
	value     int
	xValue    int
	yStart    int
	yEnd      int
	validPart bool
}

func RunPuzzle3Solution1() {
	dat, err := os.ReadFile("./puzzle3/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	fmt.Printf("Part 1: %d\n", getSumOfValidPartNos(data))

	fmt.Printf("Part 2: %d\n", getValidGearPowerSum(data))
}

func getValidGearPowerSum(data []string) int {
	// Generate a matrix of the engine schematic
	matrix := generateMatrix(data)

	// Find all the Parts represented by numbers
	parts := scanForPartNos(data)
	//fmt.Println("Parts: ", parts)

	// Find all the gear positions
	gearPositions := scanForGearPositions(matrix)
	//fmt.Println("Gears: ", gearPositions)

	// Find the valid Gear Ratios
	validGearRatios := findValidGearRatios(parts, gearPositions)

	// Sum all the valid gear ratios
	sum := calculateSumOfGearRatios(validGearRatios)

	return sum
}

func calculateSumOfGearRatios(validGearRatios []int) int {
	sum := 0
	for _, value := range validGearRatios {
		sum += value
	}
	return sum
}

func scanForGearPositions(matrix [][]string) []Position {
	//re := regexp.MustCompile(`^[a-z][A-Z][0-9].`)
	var gearPositions []Position
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			if matrix[i][j] == "*" {
				// Save the symbol position
				gearPositions = append(gearPositions, Position{i, j})
			}
		}
	}
	return gearPositions
}

func findValidGearRatios(partsMap map[int][]Parts, gearPositions []Position) []int {
	var validGearRatio []int
	for _, symbol := range gearPositions {
		// We have to check three rows of parts top, inline and bottom
		partsToCheck := partsMap[symbol.x]
		partsToCheck = append(partsToCheck, partsMap[symbol.x-1]...)
		partsToCheck = append(partsToCheck, partsMap[symbol.x+1]...)
		var validGears []Parts
		for _, part := range partsToCheck {
			// Ignore valid parts
			if part.validPart {
				continue
			}

			// Parts on top or bottom of the symbol
			if part.yStart <= symbol.y && part.yEnd >= symbol.y {
				part.validPart = true
				validGears = append(validGears, part)
				continue
			}
			// Parts left adjacent to the symbol on any three rows
			if part.yStart <= symbol.y-1 && part.yEnd == symbol.y-1 {
				part.validPart = true
				validGears = append(validGears, part)
				continue
			}
			// Parts right adjacent to the symbol on any three rows
			if part.yStart == symbol.y+1 && part.yEnd >= symbol.y+1 {
				part.validPart = true
				validGears = append(validGears, part)
				continue
			}
		}

		// Calculate the gear ratio
		if len(validGears) == 2 {
			gearRatio := validGears[0].value * validGears[1].value
			validGearRatio = append(validGearRatio, gearRatio)
		}
	}
	return validGearRatio
}

func getSumOfValidPartNos(data []string) int {
	// Generate a matrix of the engine schematic
	matrix := generateMatrix(data)

	// Find all the Parts represented by numbers
	parts := scanForPartNos(data)
	fmt.Println("Parts: ", parts)

	// Find all the symbol positions
	symbolPositions := scanForSymbolPositions(matrix)
	fmt.Println("Symbols: ", symbolPositions)

	// Find the valid parts
	validParts := findValidParts(parts, symbolPositions)

	// Sum all the valid part numbers
	sum := calculateSum(validParts)
	return sum
}

func calculateSum(validParts []Parts) int {
	sum := 0
	fmt.Println("Valid Parts: ", validParts)
	for _, part := range validParts {
		sum += part.value
	}
	return sum
}

func findValidParts(partsMap map[int][]Parts, symbolPositions []Position) []Parts {
	var validParts []Parts
	for _, symbol := range symbolPositions {
		// We have to check three rows of parts top, inline and bottom
		partsToCheck := partsMap[symbol.x]
		partsToCheck = append(partsToCheck, partsMap[symbol.x-1]...)
		partsToCheck = append(partsToCheck, partsMap[symbol.x+1]...)

		for _, part := range partsToCheck {
			// Ignore valid parts
			if part.validPart {
				continue
			}

			// Parts on top or bottom of the symbol
			if part.yStart <= symbol.y && part.yEnd >= symbol.y {
				part.validPart = true
				validParts = append(validParts, part)
				continue
			}
			// Parts left adjacent to the symbol on any three rows
			if part.yStart <= symbol.y-1 && part.yEnd == symbol.y-1 {
				part.validPart = true
				validParts = append(validParts, part)
				continue
			}
			// Parts right adjacent to the symbol on any three rows
			if part.yStart == symbol.y+1 && part.yEnd >= symbol.y+1 {
				part.validPart = true
				validParts = append(validParts, part)
				continue
			}
		}
	}
	return validParts
}

func scanForPartNos(data []string) map[int][]Parts {
	partsMap := make(map[int][]Parts, len(data))
	re := regexp.MustCompile(`\d+`)
	for i, value := range data {
		partNo := re.FindAllString(value, -1)
		// fmt.Println("Part Values: ", partNo)
		partInd := re.FindAllStringIndex(value, -1)
		// fmt.Println("Part Index: ", partInd)
		var parts []Parts
		for j, value := range partNo {
			partValue, _ := strconv.Atoi(value)
			parts = append(parts, Parts{value: partValue, xValue: i, yStart: partInd[j][0], yEnd: partInd[j][1] - 1, validPart: false})
		}
		partsMap[i] = parts
	}
	return partsMap
}

func generateMatrix(data []string) [][]string {
	fileLength := len(data)
	matrix := make([][]string, fileLength)
	for i, value := range data {
		matrix[i] = strings.Split(value, "")
	}
	return matrix
}

func scanForSymbolPositions(matrix [][]string) []Position {
	//re := regexp.MustCompile(`^[a-z][A-Z][0-9].`)
	var symbolPositions []Position
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			if matrix[i][j] != "." && !unicode.IsDigit(rune(matrix[i][j][0])) {
				// Save the symbol position
				symbolPositions = append(symbolPositions, Position{i, j})
			}
		}
	}
	return symbolPositions
}
