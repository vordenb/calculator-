package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

func Add(a, b int) (int, int) {
	total := 0
	total = a + b
	fmt.Println(total)
	return total, 0
}

func Subtract(a, b int) (int, int) {
	total := 0
	total = a - b
	fmt.Println(total)

	return total, 0
}

func Multiply(a, b int) (int, int) {
	total := 0
	total = a * b
	fmt.Println(total)

	return total, 0
}

func Divide(a, b int) (int, error) {
	total := 0
	total = a / b
	fmt.Println(total)
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	addition, _ := Add(9, 10)
	fmt.Println(addition)

	subtraction, _ := Subtract(10, 7)
	fmt.Println(subtraction)

	multiplication, _ := Multiply(2, 8)
	fmt.Println(multiplication)

	a, b := 5, 0
	result, err := Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("Divide by zero")
		default:
			fmt.Printf("unknown error: %v\n", err)
		}
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}
