package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func presentOptions() {
	fmt.Println(randomdata.SillyName())
	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
}
