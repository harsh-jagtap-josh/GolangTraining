package main

import (
	"fmt"
)

func acceptAnything(input interface{}) {

	switch valueType := input.(type) {
	case int:
		fmt.Println("This is a value of type Integer, Value is: ", valueType)
	case string:
		fmt.Println("This is a value of type String, Value is: ", valueType)
	case bool:
		fmt.Println("This is a value of type Boolean, Value is: ", valueType)
	default:
		fmt.Println("This is a value of type Hello, Value is: ", valueType)
	}
}

func main() {

	var userInput int

	fmt.Println(`Please Enter a Number 
	1: For Integer,
	2: For String,
	3: For Boolean,
	4: For any other Type`)
	fmt.Print("Enter a Number: ")
	_, err := fmt.Scanln(&userInput)

	if err != nil {
		fmt.Errorf("Error found: ", err)
	}

	var a = []int{10, 20, 30}

	switch userInput {
	case 1:
		acceptAnything(4)
	case 2:
		acceptAnything("Hello")
	case 3:
		acceptAnything(true)
	case 4:
		acceptAnything(a)
	}

}
