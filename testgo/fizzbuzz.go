package main

import "strconv"

func FizzBuzz(n int) string {
	return MappingFizzBuzz(n)
}

func MappingFizzBuzz(n int) string {
	fizz := n%3 == 0

	result := map[bool]string{
		true:  "Fizz",
		false: strconv.Itoa(n),
	}[fizz]
	return result
}
