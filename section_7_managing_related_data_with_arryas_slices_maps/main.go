package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Alex"

	userNames = append(userNames, "Maria")
	userNames = append(userNames, "John")

	fmt.Println(userNames)

}
