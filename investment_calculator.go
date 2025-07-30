package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	//pointer important in go
	var investmentAmount float64
	var years float64
	var expectedReturnRate = 5.5

	fmt.Println("|||||||||||||||||||||||||||||||||||||||||")
	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Println("|||||||||||||||||||||||||||||||||||||||||")
	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("|||||||||||||||||||||||||||||||||||||||||")
	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f", futureRealValue)

	// outputs information
	// fmt.Println("Future Valuee:", futureValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	//return
}
