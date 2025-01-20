package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string, 1)

	criticalSection := func(id int) {
		fmt.Printf("Goroutine %d is trying to access the critical section\n", id)

		ch <- "sample string"

		fmt.Printf("Goroutine %d accessing the critical section\n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Goroutine %d is leaving the critical section\n", id)
		// we basically release the section so that other routine can access it..
		fmt.Println("Critical section released: ", <-ch)
	}

	criticalSection1 := func(id int) {
		fmt.Printf("-- Goroutine %d is trying to access the critical section\n", id)

		ch <- "sample string"
		fmt.Printf("-- Goroutine %d accessing the critical section\n", id)
		time.Sleep(1 * time.Second)
		// <-ch

		// here we are not releasing the critical section to monitor the behavior of program when another routine tries to access it..
	}

	fmt.Println("Here 2 different routines try to access the same section but if one acquires the section, it releases before the other tries to access it..\n")
	for i := 1; i <= 2; i++ {
		go criticalSection(i)
	}

	time.Sleep(7 * time.Second)

	for i := 1; i <= 2; i++ {
		go criticalSection1(i)
	}

	time.Sleep(5 * time.Second)
}

// basically in the second section of the program, we observe that whichever routine gains access to the section first doesn't eventually release it,
// hence creating a problem for the second one to gain access.. Hence in the output we see the section which is first gains access but he next one eventually fails and the program is terminated as the sleep time finishes exiting out of main...
