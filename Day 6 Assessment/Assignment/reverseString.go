// 2. Given a string, reverse it using one go routine.And inside go routine print reversed string and number of goroutines launched

// e.g: Input: test123 output: 321tset 2

package main

import (
	"fmt"
	"runtime" // contains operations that interact with Go's runtime system
	"time"
)

func reverseString(str *string) {
	runes := []rune(*str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*str = string(runes)

	goRoutinesLaunched := runtime.NumGoroutine() // defines the number of Go Routines launched in a program

	fmt.Println(*str, goRoutinesLaunched)

}

func main() {

	var inputString string = "test123"

	go reverseString(&inputString)

	time.Sleep(time.Millisecond * 5) // waiting for all routines to finish
}
