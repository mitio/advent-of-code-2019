package main

import "fmt"

func isValidPassword(password int) bool {
	var digits [6]int

	for i := 5; i >= 0; i-- {
		digits[i] = password % 10
		password /= 10
	}

	atLeastTwoAdjacentDigitsAreTheSame := false
	exactlyTwoAdjacentDigitsAreTheSame := false
	sameDigits := 1

	for i := 0; i < 5; i++ {
		if digits[i] == digits[i+1] {
			atLeastTwoAdjacentDigitsAreTheSame = true
			sameDigits++
		} else {
			exactlyTwoAdjacentDigitsAreTheSame = exactlyTwoAdjacentDigitsAreTheSame || sameDigits == 2
			sameDigits = 1
		}

		if digits[i] > digits[i+1] {
			return false
		}
	}
	exactlyTwoAdjacentDigitsAreTheSame = exactlyTwoAdjacentDigitsAreTheSame || sameDigits == 2

	return atLeastTwoAdjacentDigitsAreTheSame && exactlyTwoAdjacentDigitsAreTheSame
}

func main() {
	start, end := 136818, 685979
	validPasswords := 0

	for password := start; password <= end; password++ {
		if isValidPassword(password) {
			validPasswords++
		}
	}

	fmt.Printf("valid passwords: %d\n", validPasswords)
}
