package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("Cannot divide %.2f with 0", de.dividend)
}

func Divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}

	return dividend / divisor, nil
}

func main() {
	num, err := Divide(12, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The answer is : %.2f\n", num)

}
