package puzzle6

import (
	"bufio"
	"log"
	"os"

	"github.com/golang-collections/collections/queue"
	"golang.org/x/exp/slices"
)

const packetDistinctCharCount = 4
const messageDistinctCharCount = 14

func RunPuzzle6() {
	f, err := os.Open("./puzzle6/Input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		dataStream := scan.Text()
		log.Println(dataStream)
		index := findMarker(dataStream)
		log.Println("Marker is at: ", index+1)
	}
}

func findMarker(dataStream string) int {
	slidingWindow := queue.New()
	index := 0
	for i, v := range dataStream {
		// log.Println("Index: ", i, " Value: ", v)
		index = i
		slidingWindow.Enqueue(v)
		if i < messageDistinctCharCount-1 {
			continue
		}

		if slidingWindow.Len() > messageDistinctCharCount {
			slidingWindow.Dequeue()
		}

		isUnique := checkUniqueString(slidingWindow)
		if isUnique {
			break
		}
	}
	return index
}

func checkUniqueString(slidingWindow *queue.Queue) bool {
	check := make([]rune, 0)
	l := slidingWindow.Len()
	isUnique := true
	for i := 0; i < l; i++ {
		currStr := slidingWindow.Dequeue().(rune)
		if slices.Contains(check, currStr) {
			isUnique = false
		}
		check = append(check, currStr)
		slidingWindow.Enqueue(currStr)
	}
	return isUnique
}
