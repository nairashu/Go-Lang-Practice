package loops

import (
	"fmt"
	"time"
)

func RunLoops() {

	for timer := 10; timer >= 0; timer-- {
		fmt.Println(timer)
		if timer == 0 {
			fmt.Println("BOOM!")
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("For as while condition")
	timer := 10
	for timer >= 0 {
		fmt.Println(timer)
		if timer == 0 {
			fmt.Println("BOOM!")
		}
		time.Sleep(1 * time.Second)
		timer--
	}

	fmt.Println("For as a range condition")
	courses := []string{
		"Kubernetes",
		"Docker",
		"Terra",
		"AI",
		"Machine Learning"}

	for ind, val := range courses {
		fmt.Println("Index: ", ind, " Value: ", val)
	}

	fmt.Println("Nested Loop")
	coursesCompleted := []string{
		"Docker",
		"AI"}

	for ind, val := range courses {
		fmt.Print("Index: ", ind, " Value: ", val)
		for _, j := range coursesCompleted {
			fmt.Print(" Inner Loop course ", j)
			if val == j {
				fmt.Print(" Course Complete \n")
				break
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("Break Loop")
	for boomTimer := 10; boomTimer >= 0; boomTimer-- {
		fmt.Println(boomTimer)
		if boomTimer == 5 {
			fmt.Println("Bomb defused!")
			break
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Continue Loop")
	for boomTimer := 10; boomTimer >= 0; boomTimer-- {
		if boomTimer%2 == 0 {
			continue
		}
		fmt.Println(boomTimer)
		time.Sleep(1 * time.Second)
	}
}
