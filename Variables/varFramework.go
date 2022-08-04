package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name, course string
	module, clip int

	school = "HCCS"
	rollNo = "24" //Needs to be an integer
)

func main() {
	fmt.Println("Name and Course are", name, " and ", course)
	fmt.Println("Module and Clip are", module, " and ", clip)

	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	fmt.Println("School is", school)
	fmt.Println("Roll No is", rollNo)

	fmt.Println("School is of type", reflect.TypeOf(school))
	fmt.Println("Roll No is of type", reflect.TypeOf(rollNo))

	clip = 10

	iRollNo, err := strconv.Atoi(rollNo)
	if err == nil {
		sum := clip + iRollNo
		fmt.Println("Sum is", sum)
	}

	name = "Ashish"
	course = "Kubernetes"

	fmt.Println("Name is", name)
	fmt.Println("Memory address of Course is", &course)

	var ptr *string = &course
	fmt.Println("Pointer to course is ", ptr, " and it is referencing the value ", *ptr)

	updateCourseByValue(course)
	fmt.Println("Course in main function after updateByValue is ", course)

	updateCourseByReference(&course)
	fmt.Println("Course in main function after updateByReference is ", course)
}

func updateCourseByValue(course string) string {
	course = "Docker"
	fmt.Println("Course in updateByValue function is ", course)
	return course
}

func updateCourseByReference(ptr *string) string {
	*ptr = "Docker"
	fmt.Println("Course in updateByValue function is ", *ptr)
	return *ptr
}
