package main

import (
	"fmt"
	"sync"
	"time"
)

func checkSsEven(n int) bool {
	return n%2 == 0
}

func main() {

	n := 0
	var mut sync.Mutex

	go func() {

		mut.Lock()
		nIsEven := checkSsEven(n)

		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
		mut.Unlock()

	}()

	go func() {
		mut.Lock()
		n++
		mut.Unlock()
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
