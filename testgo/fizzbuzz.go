package main

import "strconv"

func FizzBuzz(n int) string {
	if IsBuzz(n) {
		return "Buzz"
	}
	if IsFizz(n) {
		return "Fizz"
	}
	return strconv.Itoa(n)
}

func IsFizz(n int) bool {
	return n%3 == 0
}

func IsBuzz(n int) bool {
	return n%5 == 0
}
