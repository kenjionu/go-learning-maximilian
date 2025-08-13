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

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to read balance file: " + err.Error())
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse balance: " + err.Error())
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------------------------------")
		panic("Can't continue sorry.")
	}
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
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)

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
