package main

import (
	"fmt"
)

func main() {

	chan1 := make(chan int, 2)

	chan1 <- 2
	chan1 <- 1

	fmt.Println(<-chan1)
	fmt.Println(<-chan1)
	close(chan1)
	// chan1 <- 3
	fmt.Println(<-chan1)

	// go func() {
	// 	fmt.Println(<-chan1)
	// 	fmt.Println(<-chan1)
	// 	fmt.Println(<-chan1)
	// }()

	// go func() {
	// 	chan1 <- "first string"
	// 	chan1 <- "second string"
	// 	chan1 <- "third string"
	// }()

	// str1 := <-chan1
	fmt.Println("Done Executing!")
	// time.Sleep(time.Second * 1)
}
