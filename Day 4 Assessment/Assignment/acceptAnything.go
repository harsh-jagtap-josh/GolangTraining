package main

import (
	"fmt"
	"log"
)

// func acceptAnything(input interface{}){
// 	switch inputType := input.(type){
// 		case
// 	}
// }

func main() {
	// take input from user
	var input interface{}

	if _, err := fmt.Scan(&input); err != nil {
		log.Fatal("Failed to read input")
	}

	fmt.Println(input)

}
