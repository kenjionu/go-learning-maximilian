package main

import (
	"fmt"
	"time"
)

//como definir un struct type = como hacer una estructura de datos
//los structs son colecciones de campos
//los structs son tipos de datos compuestos
//los structs son tipos de datos personalizados
//los structs son tipos de datos que pueden contener otros tipos de datos

type user struct {
	firstName string
	lastName  string
	birthdate string
	createAt  time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	OutputUserDetails(lastName, firstName, birthdate)
}

func OutputUserDetails(firstName, lastName, birthdate string) {
	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
