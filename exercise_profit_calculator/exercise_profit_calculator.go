package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
// => Sow error message & exit if invalid input is provided
// Not 0
// 2) Store calculated results into file

const calculatedFile = "calculated.txt"

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("ebt: %.1f\nprofit: %.1f\nratio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(calculatedFile, []byte(results), 0644)
}

func main() {

	fmt.Println("Welcome to Profit Calculator")
	revenue, err := getUserInput("Revenue :")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------------------------------")
		panic("Can't continue sorry.")
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------------------------------")
		panic("Can't continue sorry.")
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------------------------------")
		panic("Can't continue sorry.")
	}

	ebt, profit, ratio := calculateTotally(revenue, taxRate, expenses)

	fmt.Printf("ebt: %.1f ", ebt)
	fmt.Printf("profit: %.1f ", profit)
	fmt.Printf("ratio: %.3f ", ratio)
	storeResults(ebt, profit, ratio)
}

func getUserInput(infotText string) (float64, error) {
	var userInput float64
	fmt.Print(infotText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("input must be greater than 0")
	}
	return userInput, nil
}

func calculateTotally(revenue, taxRate, expenses float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
