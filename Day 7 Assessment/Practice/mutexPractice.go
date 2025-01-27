package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

// func main() {
// 	var w sync.WaitGroup
// 	for i := 0; i < 1000; i++ {
// 		w.Add(1)
// 		go increment(&w)
// 	}
// 	w.Wait()
// 	fmt.Println("final value of x without Mutex: ", x)
// }

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementWithMutex(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x with Mutex: ", x)
}

// same can also be done using channels

func incrementWithMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
