package main

import "strconv"

func FizzBuzz(n int) string {
	result := MappingValue(n)
	return result
}

func MappingValue(n int) string {
	mapData := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	if _, ok := mapData[n]; !ok {
		return strconv.Itoa(n)
	}

	return mapData[n]
}
