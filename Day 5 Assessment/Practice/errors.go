package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Denominator cannot be zero")
	}

	return a / b, nil
}

func main() {

	a, b := 10, 0

	val, err := div(a, b)

	if err != nil {
		fmt.Println("Error Occured: ", err)
	} else {
		fmt.Println("Answer is: ", val)
	}

}
