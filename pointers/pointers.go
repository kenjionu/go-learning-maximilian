package main

import "fmt"

//go garbage collector in this case when use the function the variable age is copied to the function stack frame
// and when the function ends the stack frame is deleted and the variable age is lost
// if we want to modify the variable age in the function we need to use pointers

func main() {
	age := 32 /// regular variable

	var agePointer *int

	agePointer = &age // pointer to the variable age
	// fmt.Println("Age Pointer Value: ", *agePointer) // dereference the pointer to get the value

	fmt.Println("Age: ", *agePointer)
	//adultYears := getAdultYears(age)
	//fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
