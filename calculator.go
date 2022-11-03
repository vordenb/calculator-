package calculator

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
