package main

import "fmt"

func Luhn(cardNumber string) bool {
	var sum int
	Double := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if Double {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit

		Double = !Double
	}

	return sum%10 == 0
}

func main() {
	var cardNumber string
	fmt.Print("Enter a card number: ")
	fmt.Scan(&cardNumber)

	if Luhn(cardNumber) {
		fmt.Println("The card number is valid.")
	} else {
		fmt.Println("The card number is invalid.")
	}
}
