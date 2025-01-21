package main

import (
	"fmt"
	"sync"
)

func reverseList(list *[]string, input string, wg sync.WaitGroup) {
	for i := 0; i < len(input)/2; i++ {
		wg.Add(1)
		go func(input *[]string) {
			defer wg.Done()
			temp = list[i]
			list[i] = list[len(list)-i-1]
			list[len(list)-i-1] = temp
			fmt.Println(i, " changed")
		}(&list)
	}
}

func main() {
	input := "helloWorld"
	var list = []string{}
	var wg sync.WaitGroup
	for _, byte := range input {
		list = append(list, string(byte))
	}
	temp := ""
	fmt.Println(list)

	wg.Wait()
	fmt.Println(list)
}
