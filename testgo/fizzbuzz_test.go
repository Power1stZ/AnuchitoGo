package main

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	input := 1

	got := FizzBuzz(input)

	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn2WhenInput2(t *testing.T) {
	input := 2
	got := FizzBuzz(input)

	want := "2"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestIsFizzShouldReturnTrueWhenInput3(t *testing.T) {
	input := 3
	got := IsFizz(input)

	want := true
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestIsFizzShouldReturnFalseWhenInput4(t *testing.T) {
	input := 4
	got := IsFizz(input)

	want := false
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestIsBuzzShouldReturnTrueWhenInput5(t *testing.T) {
	input := 5
	got := IsBuzz(input)

	want := true
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput6(t *testing.T) {
	input := 6
	got := FizzBuzz(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestFizzBuzzShouldReturn7WhenInput7(t *testing.T) {
	input := 7
	got := FizzBuzz(input)

	want := "7"
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestFizzBuzzShouldReturn8WhenInput8(t *testing.T) {
	input := 8
	got := FizzBuzz(input)

	want := "8"
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput9(t *testing.T) {
	input := 9
	got := FizzBuzz(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestFizzBuzzShouldReturnBuzzWhenInput10(t *testing.T) {
	input := 10
	got := FizzBuzz(input)

	want := "Buzz"
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestFizzBuzzShouldReturn11WhenInput11(t *testing.T) {
	input := 11
	got := FizzBuzz(input)

	want := "11"
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput12(t *testing.T) {
	input := 12
	got := FizzBuzz(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestFizzBuzzShouldReturn13WhenInput13(t *testing.T) {
	input := 13
	got := FizzBuzz(input)

	want := "13"
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
