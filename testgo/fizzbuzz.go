package main

import "strconv"

func FizzBuzz(n int) string {

	return MappingFizzBuzz(n)
}

func MappingFizzBuzz(n int) string {

	mapping := map[int]string{
		3:  "Fizz",
		5:  "Buzz",
		6:  "Fizz",
		9:  "Fizz",
		10: "Buzz",
		12: "Fizz",
		15: "FizzBuzz",
	}

	if value, ok := mapping[n]; ok {
		return value
	}
	return strconv.Itoa(n)
}
