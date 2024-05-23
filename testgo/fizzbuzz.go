package main

import "strconv"

func FizzBuzz(n int) string {
	return MappingFizzBuzz(n)
}

func MappingFizzBuzz(n int) string {
	fizz := n%3 == 0
	buzz := n%5 == 0

	result := map[bool]string{
		true:  "Fizz" + map[bool]string{true: "Buzz", false: ""}[buzz],
		false: map[bool]string{true: "Buzz", false: strconv.Itoa(n)}[buzz],
	}[fizz]
	return result
}
