package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//instance of a struct
	var appUser *user.User

	//using a constructor function to create a new user
	appUser, err := &user.User{
		FirstName: "Max",
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	//pay attention to the order of the fields
	appUser = &user{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}
	// ... do something awesome with that gathered data!

	//accesing struct fields with pointer
	//OutputUserDetails(&appUser)

	//accesing struct fields without pointer
	//OutputUserDetails(appUser)

	//accesing struct fields with method
	appUser.OutputUserDetails()
	appUser.clearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
