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
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//instance of a struct
	var appUser user

	//pay attention to the order of the fields
	appUser = user{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}
	// ... do something awesome with that gathered data!

	//accesing struct fields with pointer
	OutputUserDetails(&appUser)
}

// add pointer parameter to avoid copying the struct
func OutputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
