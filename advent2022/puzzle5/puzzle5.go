package puzzle5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func RunPuzzle5() {
	f, err := os.Open("./puzzle5/Test1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	readCrates := true
	stacksOfCrate := make([]stack.Stack, 0)

	for scanner.Scan() {
		scanLine := scanner.Text()

		// Create the stacks with the information of the crates from the diagram
		if readCrates {
			crateStrings := getCrateStrings(scanLine)
			log.Println(crateStrings)
			if strings.Contains(crateStrings[0], "1") {
				readCrates = false
				continue
			}
			stacksOfCrate = updateStacksWithCrateStrings(stacksOfCrate, crateStrings)
			log.Println(len(stacksOfCrate))
			continue
		}

		// Read all the movement information and move the stack accordingly
		if !strings.Contains(scanLine, "move") {
			stacksOfCrate = reverseStacks(stacksOfCrate)
			continue
		}

		moves := strings.Split(scanLine, " ")
		crateCount, _ := strconv.Atoi(moves[1])
		srcStack, _ := strconv.Atoi(moves[3])
		dstStack, _ := strconv.Atoi(moves[5])
		stacksOfCrate = performAction(stacksOfCrate, crateCount, srcStack-1, dstStack-1)
	}
	topOfStacks := getTopOfStacks(stacksOfCrate)
	log.Println(topOfStacks)
}

func reverseStacks(stacksOfCrate []stack.Stack) []stack.Stack {
	fixedCrateStack := make([]stack.Stack, 0)
	for i := 0; i < len(stacksOfCrate); i++ {
		fixedCrateStack = append(stacksOfCrate, *stack.New())

		for j := 0; j < stacksOfCrate[i].Len(); j++ {
			item := stacksOfCrate[i].Pop().(string)
			fixedCrateStack[i].Push(item)
		}
	}
	return fixedCrateStack
}

func getTopOfStacks(stacksOfCrate []stack.Stack) string {
	var builder strings.Builder
	for i := 0; i < len(stacksOfCrate); i++ {
		t := stacksOfCrate[i].Peek().(string)
		t = strings.ReplaceAll(t, "[", "")
		t = strings.ReplaceAll(t, "]", "")
		_, err := builder.WriteString(t)
		if err != nil {
			log.Fatal(err)
		}
	}
	return strings.TrimSpace(builder.String())
}

func performAction(stacksOfCrate []stack.Stack, crateCount, srcStack, dstStack int) []stack.Stack {
	for i := 0; i < crateCount; i++ {
		crate := stacksOfCrate[srcStack].Pop()
		stacksOfCrate[dstStack].Push(crate)
	}
	return stacksOfCrate
}

func updateStacksWithCrateStrings(stacksOfCrate []stack.Stack, crateStrings []string) []stack.Stack {
	for i, v := range crateStrings {
		if len(stacksOfCrate) <= i {
			stacksOfCrate = append(stacksOfCrate, *stack.New())
		}

		if !strings.Contains(v, " ") {
			stacksOfCrate[i].Push(v)
		}
	}
	return stacksOfCrate
}

func getCrateStrings(scanLine string) []string {
	stackDetails := []rune(scanLine)
	crates := make([]string, len(stackDetails)/3)
	for i := 0; i < len(stackDetails); i = i + 4 {
		crates[i/3] = string([]rune{stackDetails[i], stackDetails[i+1], stackDetails[i+2]})
	}
	return crates
}
