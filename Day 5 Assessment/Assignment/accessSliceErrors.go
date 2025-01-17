package main

import (
	"errors"
	"fmt"
)

func accessSliceAgain(array []int, index int) (string, error) {
	if index >= len(array) {
		return "", errors.New("Lenght of slice should be more than index.")
	}

	str := fmt.Sprintf("Item: %v, Value: %v", index, array[index])

	return str, nil
}

func main() {
	var array = []int{10, 20, 30, 40, 50, 60}
	var index int
	fmt.Printf("Enter the index for an array of length %v: ", len(array))
	fmt.Scan(&index)
	val, err := accessSliceAgain(array, index)

	if err != nil {
		fmt.Println("Error occured: ", err)
	}

	fmt.Println(val)
}
