package puzzle5

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var Seeds []int
var SoilBySeed = make(map[int]int)
var FertilizerBySoil = make(map[int]int)
var WaterByFertilizer = make(map[int]int)
var LightByWater = make(map[int]int)
var TemperatureByLight = make(map[int]int)
var HumidityByTemperature = make(map[int]int)
var LocationByHumidity = make(map[int]int)

func RunPuzzle5Solution1() {
	// dat, err := os.ReadFile("./puzzle5/Example1.txt")
	dat, err := os.ReadFile("./puzzle5/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	fmt.Printf("Part 1: %d\n", getShortestLocation(data))

	// fmt.Printf("Part 2: %d\n", calculateSumOfAllCopiesOfScratchcards(data))
}

func getShortestLocation(data []string) int {
	readDataAndInitializeMaps(data)
	// fmt.Println("Seeds: ", Seeds)
	// fmt.Println("SoilBySeed: ", SoilBySeed)
	// fmt.Println("FertilizerBySoil: ", FertilizerBySoil)
	// fmt.Println("WaterByFertilizer: ", WaterByFertilizer)
	// fmt.Println("LightByWater: ", LightByWater)
	// fmt.Println("TemperatureByLight: ", TemperatureByLight)
	// fmt.Println("HumidityByTemperature: ", HumidityByTemperature)
	// fmt.Println("LocationByHumidity: ", LocationByHumidity)
	location := findShortestLocation()
	return location
}

func findShortestLocation() int {
	locations := make([]int, len(Seeds))
	for i, seed := range Seeds {
		soil, ok := SoilBySeed[seed]
		if !ok {
			soil = seed
		}
		fertilizer, ok := FertilizerBySoil[soil]
		if !ok {
			fertilizer = soil
		}
		water, ok := WaterByFertilizer[fertilizer]
		if !ok {
			water = fertilizer
		}
		light, ok := LightByWater[water]
		if !ok {
			light = water
		}
		temperature, ok := TemperatureByLight[light]
		if !ok {
			temperature = light
		}
		humidity, ok := HumidityByTemperature[temperature]
		if !ok {
			humidity = temperature
		}
		location, ok := LocationByHumidity[humidity]
		if !ok {
			location = humidity
		}
		locations[i] = location
		fmt.Printf("Seed: %d, Soil: %d, Fertilizer: %d, Water: %d, Light: %d, Temperature: %d, Humidity: %d, Location: %d\n", seed, soil, fertilizer, water, light, temperature, humidity, location)
	}
	// fmt.Println("Locations: ", locations)
	sort.Ints(locations)
	return locations[0]
}

func readDataAndInitializeMaps(data []string) {
	seedsloaded := false
	loadSoil := false
	loadFertilizer := false
	loadWater := false
	loadLight := false
	loadTemperature := false
	loadHumidity := false
	loadLocation := false

	for _, line := range data {
		re := regexp.MustCompile(`\d+`)
		if strings.Contains(line, "seeds") && !seedsloaded {
			lineData := strings.Split(line, ":")
			seeds := re.FindAllString(lineData[1], -1)
			convertAndLoadSeedsSlice(seeds)
			seedsloaded = true
			continue
		}

		if len(line) == 0 {
			loadSoil = false
			loadFertilizer = false
			loadWater = false
			loadLight = false
			loadTemperature = false
			loadHumidity = false
			loadLocation = false
			continue
		}

		if strings.Contains(line, "seed-to-soil map") {
			loadSoil = true
			continue
		} else if strings.Contains(line, "soil-to-fertilizer map") {
			loadFertilizer = true
			continue
		} else if strings.Contains(line, "fertilizer-to-water map") {
			loadWater = true
			continue
		} else if strings.Contains(line, "water-to-light map") {
			loadLight = true
			continue
		} else if strings.Contains(line, "light-to-temperature map") {
			loadTemperature = true
			continue
		} else if strings.Contains(line, "temperature-to-humidity map") {
			loadHumidity = true
			continue
		} else if strings.Contains(line, "humidity-to-location map") {
			loadLocation = true
			continue
		}

		if loadSoil {
			loadDataIntoMap(line, "SoilBySeed")
		} else if loadFertilizer {
			loadDataIntoMap(line, "FertilizerBySoil")
		} else if loadWater {
			loadDataIntoMap(line, "WaterByFertilizer")
		} else if loadLight {
			loadDataIntoMap(line, "LightByWater")
		} else if loadTemperature {
			loadDataIntoMap(line, "TemperatureByLight")
		} else if loadHumidity {
			loadDataIntoMap(line, "HumidityByTemperature")
		} else if loadLocation {
			loadDataIntoMap(line, "LocationByHumidity")
		}
	}
}

func loadDataIntoMap(line, s string) {
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(line, -1)
	if len(numbers) != 3 {
		log.Fatal("Invalid data")
	}
	// fmt.Println(s, " Numbers: ", numbers)
	destValue, _ := strconv.Atoi(numbers[0])
	srcKey, _ := strconv.Atoi(numbers[1])
	rangeNumber, _ := strconv.Atoi(numbers[2])
	for i := 0; i <= rangeNumber; i++ {
		switch s {
		case "SoilBySeed":
			SoilBySeed[srcKey+i] = destValue + i
		case "FertilizerBySoil":
			FertilizerBySoil[srcKey+i] = destValue + i
		case "WaterByFertilizer":
			WaterByFertilizer[srcKey+i] = destValue + i
		case "LightByWater":
			LightByWater[srcKey+i] = destValue + i
		case "TemperatureByLight":
			TemperatureByLight[srcKey+i] = destValue + i
		case "HumidityByTemperature":
			HumidityByTemperature[srcKey+i] = destValue + i
		case "LocationByHumidity":
			LocationByHumidity[srcKey+i] = destValue + i
		}
	}
}

func convertAndLoadSeedsSlice(seeds []string) {
	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		Seeds = append(Seeds, seedInt)
	}
}

type Alamac struct {
	destValue   int
	srcKey      int
	rangeNumber int
}

var SeedToSoil []Alamac
var SoilToFertilizer []Alamac
var FertilizerToWater []Alamac
var WaterToLight []Alamac
var LightToTemperature []Alamac
var TemperatureToHumidity []Alamac
var HumidityToLocation []Alamac

func RunPuzzle5Solution2() {
	// dat, err := os.ReadFile("./puzzle5/Example1.txt")
	dat, err := os.ReadFile("./puzzle5/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	// fmt.Printf("Part 1: %d\n", getSmallestLocation(data))

	fmt.Printf("Part 2: %d\n", getSmallestLocationConsideringSeedRanges(data))
}

func getSmallestLocationConsideringSeedRanges(data []string) int {
	readDataAndInitializeLists(data, true)
	location := identifySmallestSeedLocations(Seeds)
	return location
}

func getSmallestLocation(data []string) int {
	readDataAndInitializeLists(data, false)
	location := identifySmallestSeedLocations(Seeds)
	return location
}

func identifySmallestSeedLocations(Seeds []int) int {
	locations := make([]int, len(Seeds))
	for i, seed := range Seeds {
		soil := findTargetVallue(SeedToSoil, seed)
		fertilizer := findTargetVallue(SoilToFertilizer, soil)
		water := findTargetVallue(FertilizerToWater, fertilizer)
		light := findTargetVallue(WaterToLight, water)
		temperature := findTargetVallue(LightToTemperature, light)
		humidity := findTargetVallue(TemperatureToHumidity, temperature)
		location := findTargetVallue(HumidityToLocation, humidity)
		locations[i] = location
	}
	sort.Ints(locations)
	fmt.Println("Locations: ", locations)
	return locations[0]
}

func findTargetVallue(alamacSlice []Alamac, srcTarget int) int {
	destValue := -1
	for _, alamac := range alamacSlice {
		if srcTarget < alamac.srcKey || srcTarget > alamac.srcKey+alamac.rangeNumber {
			continue
		}
		diff := srcTarget - alamac.srcKey
		destValue = alamac.destValue + diff
	}
	if destValue == -1 {
		return srcTarget
	}
	return destValue
}

func readDataAndInitializeLists(data []string, considerSeedRanges bool) {
	seedsloaded := false
	loadSoil := false
	loadFertilizer := false
	loadWater := false
	loadLight := false
	loadTemperature := false
	loadHumidity := false
	loadLocation := false

	SeedToSoil = make([]Alamac, 0)
	SoilToFertilizer = make([]Alamac, 0)
	FertilizerToWater = make([]Alamac, 0)
	WaterToLight = make([]Alamac, 0)
	LightToTemperature = make([]Alamac, 0)
	TemperatureToHumidity = make([]Alamac, 0)
	HumidityToLocation = make([]Alamac, 0)

	for _, line := range data {
		strings.Trim(line, " ")
		// fmt.Println("Line: ", line)
		re := regexp.MustCompile(`\d+`)
		if strings.Contains(line, "seeds") && !seedsloaded {
			lineData := strings.Split(line, ":")
			seeds := re.FindAllString(lineData[1], -1)
			convertAndLoadSeeds(seeds, considerSeedRanges)
			seedsloaded = true
			continue
		}

		if strings.Contains(line, "seed-to-soil map") {
			loadSoil = true
			continue
		} else if strings.Contains(line, "soil-to-fertilizer map") {
			loadFertilizer = true
			continue
		} else if strings.Contains(line, "fertilizer-to-water map") {
			loadWater = true
			continue
		} else if strings.Contains(line, "water-to-light map") {
			loadLight = true
			continue
		} else if strings.Contains(line, "light-to-temperature map") {
			loadTemperature = true
			continue
		} else if strings.Contains(line, "temperature-to-humidity map") {
			loadHumidity = true
			continue
		} else if strings.Contains(line, "humidity-to-location map") {
			loadLocation = true
			continue
		}

		numbers := re.FindAllString(line, -1)
		if len(numbers) == 0 {
			// fmt.Println("Empty line")
			loadSoil = false
			loadFertilizer = false
			loadWater = false
			loadLight = false
			loadTemperature = false
			loadHumidity = false
			loadLocation = false
			continue
		}

		if loadSoil {
			SeedToSoil = loadDataIntoSlice(line, SeedToSoil)
		} else if loadFertilizer {
			SoilToFertilizer = loadDataIntoSlice(line, SoilToFertilizer)
		} else if loadWater {
			FertilizerToWater = loadDataIntoSlice(line, FertilizerToWater)
		} else if loadLight {
			WaterToLight = loadDataIntoSlice(line, WaterToLight)
		} else if loadTemperature {
			LightToTemperature = loadDataIntoSlice(line, LightToTemperature)
		} else if loadHumidity {
			TemperatureToHumidity = loadDataIntoSlice(line, TemperatureToHumidity)
		} else if loadLocation {
			HumidityToLocation = loadDataIntoSlice(line, HumidityToLocation)
		}
	}
}

func convertAndLoadSeeds(seeds []string, considerSeedRanges bool) {
	if !considerSeedRanges {
		for _, seed := range seeds {
			seedInt, _ := strconv.Atoi(seed)
			Seeds = append(Seeds, seedInt)
		}
	} else {
		for i := 0; i < len(seeds); i += 2 {
			seedInt, _ := strconv.Atoi(seeds[i])
			rangeInt, _ := strconv.Atoi(seeds[i+1])
			for j := 0; j < rangeInt; j++ {
				Seeds = append(Seeds, seedInt+j)
			}
		}
	}
}

func loadDataIntoSlice(line string, slice []Alamac) []Alamac {
	// fmt.Println("Line: ", line)
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(line, -1)
	if len(numbers) != 3 {
		log.Fatalf("Invalid data %v", line)
	}
	// fmt.Println(s, " Numbers: ", numbers)
	destValue, _ := strconv.Atoi(numbers[0])
	srcKey, _ := strconv.Atoi(numbers[1])
	rangeNumber, _ := strconv.Atoi(numbers[2])
	alamac := Alamac{destValue: destValue, srcKey: srcKey, rangeNumber: rangeNumber}
	slice = append(slice, alamac)
	return slice
}
