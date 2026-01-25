package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Alex"

	userNames = append(userNames, "Maria")
	userNames = append(userNames, "John")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.8

	courseRatings.output()

	delete(courseRatings, "angular")

	fmt.Println(courseRatings)
	/// ....
	for index, value := range userNames {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	for key, value := range courseRatings {
		fmt.Printf("Key: %s, Value: %.2f\n", key, value)
	}
}
