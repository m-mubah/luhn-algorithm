package services

import (
	"log"
	"strconv"
)

// ValidateCard - if the card is valid, returns true and the checksum of a given card number.
func ValidateCard(cardNumber string) (bool, int) {
	e, err := strconv.Atoi(cardNumber[len(cardNumber)-1:])
	if err != nil {
		log.Panic(err)
	}

	a := cardNumber[:len(cardNumber)-1]

	var check = calculate(a)

	if check == e {
		return true, check
	}

	return false, 0
}

// calculate - performs luhn check for a card number
func calculate(cardNumber string) int {
	temp := 0

	for i := 0; i < len(cardNumber); i++ {
		conv, err := strconv.Atoi(cardNumber[i : i+1])
		if err != nil {
			log.Fatal(err)
		}
		temp += conv
	}

	n := []int{0, 1, 2, 3, 4, -4, -3, -2, -1, 0}

	for i := len(cardNumber) - 1; i >= 0; i -= 2 {
		conv, err := strconv.Atoi(cardNumber[i : i+1])
		if err != nil {
			log.Fatal(err)
		}
		r := n[conv]

		temp += r
	}

	s := temp % 10
	s = 10 - s

	if s == 10 {
		s = 0
	}

	return s
}
