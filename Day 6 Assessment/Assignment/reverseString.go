package main

import (
	"fmt"
	"sync"
)

func reverseString(str string) string {
	newStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		newStr = newStr + str[i:i+1]
	}

	return newStr
}

func main() {

	var inputString string = "test123"
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		// Reverse the string
		inputString = reverseString(inputString)
	}()

	wg.Wait()

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println(inputString, 2)
	}()

	wg.Wait()
}
