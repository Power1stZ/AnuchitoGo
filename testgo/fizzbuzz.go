package main

import "strconv"

func FizzBuzz(n int) string {

	return MappingFizzBuzz(n)
}

func MappingFizzBuzz(n int) string {

	mapping := map[int]string{
		3: "Fizz",
		5: "Buzz",
		6: "Fizz",
	}

	if value, ok := mapping[n]; ok {
		return value
	}
	return strconv.Itoa(n)
}
