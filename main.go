package calculator

import (
	"errors"
	"fmt"
)

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
