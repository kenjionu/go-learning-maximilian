package main

import (
	"fmt"
	"math"
)

func main() {
	//pointer important in go
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Println("|||||||||||||||||||||||||||||||||||||||||")
	fmt.Print("Investment Amount: ")

	fmt.Scan(&investmentAmount)

	fmt.Println("|||||||||||||||||||||||||||||||||||||||||")
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("|||||||||||||||||||||||||||||||||||||||||")
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f", futureRealValue)

	// outputs information
	// fmt.Println("Future Valuee:", futureValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}
