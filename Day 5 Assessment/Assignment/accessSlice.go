package main

import "fmt"

func accessSlice(array []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Internal error: ", r)
		}
	}()
	fmt.Printf("Item: %d, Value: %d", index, array[index])
	fmt.Println("Testing Panic and Recovery")
}

func main() {
	var array = []int{10, 20, 30, 40, 50, 60}
	var index int
	fmt.Printf("Enter the index for an array of length %v: ", len(array))
	fmt.Scan(&index)

	accessSlice(array, index)
}
