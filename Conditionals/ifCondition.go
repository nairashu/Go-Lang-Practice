package main

import (
	"fmt"
	"os"
)

func main() {
	dockerCourseLength := 235
	kubernetesCourseLength := 30

	if dockerCourseLength > kubernetesCourseLength {
		fmt.Println("Docker is a longer course")
	} else if dockerCourseLength == kubernetesCourseLength {
		fmt.Println("Both courses have the same duration")
	} else {
		fmt.Println("Kubernetes is a longer course")
	}

	if dockerCourseLength, kubernetesCourseLength = 20, 120; dockerCourseLength > kubernetesCourseLength {
		fmt.Println("Docker is a longer course")
	} else if dockerCourseLength == kubernetesCourseLength {
		fmt.Println("Both courses have the same duration")
	} else {
		fmt.Println("Kubernetes is a longer course")
		if kubernetesCourseLength < 150 {
			fmt.Println("The Kubernetes course is still within the viewing time limit")
		}
	}

	_, err := os.Open("./Test.txt")

	if err != nil {
		fmt.Print("The error code is ", err)
	}
}
