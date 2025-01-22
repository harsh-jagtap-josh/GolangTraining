package main

import (
	"fmt"
	"sync"
)

func convertToString(list []string) string {
	final := ""

	for _, byte := range list {
		final = final + string(byte)
	}
	return final
}

func main() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scan(&input)

	var list = []string{}
	var wg sync.WaitGroup
	for _, byte := range input {
		list = append(list, string(byte))
	}
	temp := ""

	for i := 0; i < len(input)/2; i++ {
		wg.Add(1)
		go func(input *[]string) {
			defer wg.Done()

			temp = list[i]
			list[i] = list[len(list)-i-1]
			list[len(list)-i-1] = temp
		}(&list)
	}
	wg.Wait()

	fmt.Println("Reversed String: ", convertToString(list))
}
