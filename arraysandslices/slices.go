package arraysandslices

import (
	"fmt"
)

func RunArraysAndSlices() {
	fmt.Println("Working with slices")

	courses := make([]string, 5, 10)
	courses[0] = "Docker"
	courses[1] = "Kubernetes"
	courses[2] = "AI"
	courses[3] = "Terra"

	fmt.Println("Length of courses slice is ", len(courses), " and capacity is ", cap(courses))

	for _, i := range courses {
		fmt.Println(i)
	}

	names := []string{"Ashish", "Tim", "Sujay"}

	fmt.Println("Length of names slice is ", len(names), " and capacity is ", cap(names))

	for _, i := range names {
		fmt.Println(i)
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)
	fmt.Println(numbers[4])

	numbers[1] = 0
	fmt.Println(numbers)

	subSlice := numbers[3:6]
	fmt.Println(subSlice)

	names = append(names, "Noah")
	fmt.Println(names)

	appendSlice := make([]int, 1, 4)
	for i := 1; i < 17; i++ {
		appendSlice = append(appendSlice, i)
		fmt.Println(appendSlice)
		fmt.Printf("Lenght is %d and Capacity is %d \n", len(appendSlice), cap(appendSlice))
	}

	numbers = append(numbers, appendSlice...)
	fmt.Println(numbers)
}
