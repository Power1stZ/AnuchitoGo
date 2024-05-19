package main

import "strconv"

func FizzBuzz(n int) string {
	if IsFizz(n) {
		return "Fizz"
	}

	if IsBuzz(n) {
		return "Buzz"
	}

	return strconv.Itoa(n)
}

func IsFizz(n int) bool {
	return n%3 == 0
}

func IsBuzz(n int) bool {
	return n%5 == 0
}
