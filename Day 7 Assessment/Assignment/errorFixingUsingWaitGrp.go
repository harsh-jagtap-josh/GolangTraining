package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {

	var wg sync.WaitGroup
	n := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		nIsEven := isEven(n)
		fmt.Println(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		n++
	}()
	wg.Done()
	// just waiting for the goroutines to finish before exiting
	// time.Sleep(time.Second)
}
