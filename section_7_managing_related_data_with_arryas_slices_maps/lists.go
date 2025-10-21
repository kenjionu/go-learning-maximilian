package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)

}

// func main() {
// 	var productNames [4]string = [4]string{"Laptop", "Phone", "Tablet", "Monitor"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "Smartwatch"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlitedPrices := featuredPrices[:1]
// 	fmt.Println(highlitedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlitedPrices), cap(highlitedPrices))

// 	highlitedPrices = highlitedPrices[:3]
// 	fmt.Println(highlitedPrices)
// 	fmt.Println(len(highlitedPrices), cap(highlitedPrices))

// }
