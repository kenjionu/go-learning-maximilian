package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//instance of a struct
	var appUser *user.User

	//using a constructor function to create a new user
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "supersecret")

	admin.User.OutputUserDetails()
	admin.User.ClearUserName()
	admin.User.OutputUserDetails()

	// ... do something awesome with that gathered data!

	//accesing struct fields with pointer
	//OutputUserDetails(&appUser)

	//accesing struct fields without pointer
	//OutputUserDetails(appUser)

	//accesing struct fields with method
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
