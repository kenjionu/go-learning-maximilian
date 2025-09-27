package user

import (
	"errors"
	"fmt"
	"time"
)

//como definir un struct type = como hacer una estructura de datos
//los structs son colecciones de campos
//los structs son tipos de datos compuestos
//los structs son tipos de datos personalizados
//los structs son tipos de datos que pueden contener otros tipos de datos

type User struct {
	firstName string
	lastName  string
	birthdate string
	createAt  time.Time
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
	u.createAt = time.Time{}
}

// add pointer parameter to avoid copying the struct
// method with a receiver of type user (pointer) with the struct user
// el func (u user) es el receiver

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// constructor function to create a new user
// convention to start with "new" + struct name
// returns a pointer to the struct
// helps to avoid copying the struct
// initializes the struct with the provided values
// sets the createAt field to the current time
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate cannot be empty")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createAt:  time.Now(),
	}, nil
}
