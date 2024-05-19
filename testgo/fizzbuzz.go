package main

import "strconv"

func FizzBuzz(n int) string {
	if IsFizz(n) {
		return "Fizz"
	}

	return strconv.Itoa(n)
}

func IsFizz(n int) bool {
	return n%3 == 0
}
