package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	var accountBalance float64 = 1000.00

	fmt.Println("Welcome to go Bank")
	img, err := loadIamge("bank4.png")
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
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	// wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance  is: ", accountBalance)
	} else if choice == 2 {
		fmt.Println("Your deposit")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Deposit amount must be greater than 0.")
			return
		}
		accountBalance = deposit(accountBalance, depositAmount)
		fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
	} else if choice == 3 {
		fmt.Println("Your withdraw")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)

		if withdrawalAmount <= 0 {
			fmt.Println("Withdrawal amount must be greater than 0.")
			return
		}

		if withdrawalAmount > accountBalance {
			fmt.Println("Insufficient funds: you cannot withdraw more than your current balance.")
			return
		}

		accountBalance += withdrawalAmount
		fmt.Println("Your new balance is: ", accountBalance)
	} else {
		fmt.Println("Exiting the bank. Goodbye!")
		return
	}

	fmt.Println("Your choice is: ", choice)
}

func checkBalance(balance float64) {
	fmt.Scan()
	fmt.Printf("Your current balance is: %.2f\n", balance)
}
func deposit(balance float64, amount float64) float64 {
	fmt.Printf("Depositiing ^%.2f\n", amount)
	return balance + amount
}

func withdraw(balance, amount float64) (float64, error) {
	if amount > balance {
		return balance, fmt.Errorf("Insufficient funds: you cannot withdraw %.2f from a balance of %.2f", amount, balance)
	} else {
		fmt.Printf("Withdrawing %.2f\n", amount)
		return balance - amount, nil
	}
}

func loadIamge(filePath string) (image.Image, error) {
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
