package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to read balance file: " + err.Error())
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse balance: " + err.Error())
	}

	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile)

	fmt.Println("Welcome to go Bank")
	img, err := loadImage("bank4.png")
	if err != nil {
		panic((err))
	}
	ramp := "@#+=. "
	max := img.Bounds().Max
	scaleX, scaleY := 10, 5
	for y := 0; y < max.Y; y += scaleY {
		for x := 0; x < max.X; x += scaleX {
			avg := avgPixel(img, x, y, scaleX, scaleY)
			fmt.Print(string(ramp[avg*len(ramp)/65536]))
		}
		fmt.Println()
	}
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance  is: ", accountBalance)
		case 2:
			fmt.Println("Your deposit")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
			writeFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Println("Your withdraw")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Withdrawal amount must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds: you cannot withdraw more than your current balance.")
				continue
			}

			accountBalance -= withdrawalAmount

			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Exiting the bank. Goodbye!")
			fmt.Println("Thank you for choosing go Bank!")
			return
		}
	}

}

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}

func grayscale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	gray := (r + g + b) / 3
	return int(gray)
}

func avgPixel(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += grayscale(img.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}
