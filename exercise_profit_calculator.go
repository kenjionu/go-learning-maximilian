package main

import (
	"fmt"
)

func tested() {
	var revenue float64
	var taxRate float64
	var expenses float64

	fmt.Println("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Println("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Println("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
