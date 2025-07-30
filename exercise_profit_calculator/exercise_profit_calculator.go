package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Revenue :")

	expenses := getUserInput("Expenses: ")

	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateTotally(revenue, taxRate, expenses)

	fmt.Printf("%.1f", ebt)
	fmt.Printf("%.1f", profit)
	fmt.Printf("%.3f", ratio)
}

func getUserInput(infotText string) float64 {
	var userInput float64
	fmt.Print(infotText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateTotally(revenue, taxRate, expenses float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
