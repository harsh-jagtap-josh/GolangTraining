package main

import (
	"fmt"
	"math/rand"
	"time"
)

//
// --> Send and Receive in a channel are blocking calls, they block the program until and unless the data is sent or received..
// Types of CHANNELS:
// --> Buffered and Un-buffered
// --> Unidirectional and Bidirectional

// ========================================================================================================================

// func main() {

// 	chan1 := make(chan int, 2)

// 	chan1 <- 2
// 	chan1 <- 1

// 	fmt.Println(<-chan1)
// 	fmt.Println(<-chan1)
// 	close(chan1)
// 	chan1 <- 3
// 	fmt.Println(<-chan1)
// }

// ========================================================================================================================

// func hello(done chan bool) {
// 	fmt.Println("Hello world goroutine")
// 	done <- true
// }
// func main() {
// 	done := make(chan bool)
// 	go hello(done)
// 	<-done
// 	fmt.Println("main function")
// }

// ========================================================================================================================

// func hello(done chan bool) {
// 	fmt.Println("hello go routine is going to sleep")
// 	time.Sleep(4 * time.Second)
// 	fmt.Println("hello go routine awake and going to write to done")
// 	done <- true
// }
// func main() {
// 	done := make(chan bool)
// 	fmt.Println("Main going to call hello go goroutine")
// 	go hello(done)
// 	<-done
// 	fmt.Println("Main received data")
// }

// ========================================================================================================================

// func calcSquares(number int, squareop chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit
// 		number /= 10
// 	}
// 	squareop <- sum
// }

// func calcCubes(number int, cubeop chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit * digit
// 		number /= 10
// 	}
// 	cubeop <- sum
// }

// func main() {
// 	number := 589
// 	sqrch := make(chan int)
// 	cubech := make(chan int)
// 	go calcSquares(number, sqrch)
// 	go calcCubes(number, cubech)
// 	squares, cubes := <-sqrch, <-cubech
// 	fmt.Println("Final output, Squares: ", squares, " Cubes: ", cubes)
// }

// ========================================================================================================================

// Unidirectional Channels

// func sendData(sendch chan<- int) {
// 	sendch <- 10
// }

// func main() {
// 	sendch := make(chan<- int) // a unidirectional channel that can only receive and return an error if we try to retrieve the value
// 	go sendData(sendch)
// 	fmt.Println(<-sendch)
// }

// ========================================================================================================================

// func addToChan(ch1 chan int) {
// 	for i := 0; i <= 10; i++ {
// 		ch1 <- i
// 	}

// 	close(ch1)
// }

// func main() {
// 	var newCh = make(chan int)
// 	go addToChan(newCh)

// 	for {
// 		val, ok := <-newCh
// 		if !ok {
// 			break
// 		} else {
// 			fmt.Println(val, " received from channel")
// 		}
// 	}

// 	var newChannel = make(chan int)

// 	go addToChan(newChannel)
// 	for v := range newChannel {
// 		fmt.Println(v, " received from channel in second for loop")
// 	}
// 	// time.Sleep(time.Second * 1)
// }

// ========================================================================================================================

// func main() {
// 	car1 := "Ferrari"
// 	car2 := "Lamborghini"

// 	// Create a Goroutine for each car
// 	go race(car1)
// 	go race(car2)

// 	// Wait for the race to finish
// 	time.Sleep(5 * time.Second)

// 	fmt.Println("Race over!")
// }

// func race(car string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(car, i, "is racing...")
// 		time.Sleep(1 * time.Second)
// 	}
// }

// =====================================================================================================================================

// func main() {
// 	// Create a buffered channel with a capacity of 2
// 	cars := make(chan string, 2)

// 	// Start two Goroutines that add cars to the channel
// 	go addCar("Ferrari", cars)
// 	go addCar("Lamborghini", cars)

// 	// Wait for the cars to be added to the channel
// 	time.Sleep(2 * time.Second)
// 	close(cars)

// 	// Start a Goroutine that simulates the race
// 	go startRace(cars)

// 	// Wait for the race to finish
// 	time.Sleep(6 * time.Second)

// 	fmt.Println("Race over!")
// }

// func addCar(name string, cars chan string) {
// 	cars <- name
// 	fmt.Println(name, "added to the race!")
// }

// func startRace(cars chan string) {
// 	for {
// 		// Receive a car from the channel
// 		car, open := <-cars
// 		if !open {
// 			break
// 		}

// 		fmt.Println(car, "is racing...")

// 		// Simulate the race by waiting for a random duration
// 		time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
// 	}

// 	fmt.Println("All cars have finished the race!")
// }

// =====================================================================================================================================

func main() {
	// Create two channels for the cars
	ferrari, lamborghini := make(chan string), make(chan string)

	// Start two Goroutines that add cars to the channels
	go addCarWithChan("Ferrari", ferrari)
	go addCarWithChan("Lamborghini", lamborghini)

	// Start a Goroutine that simulates the race
	go startRaceWithSelect(ferrari, lamborghini)

	// Wait for the race to finish
	time.Sleep(5 * time.Second)

	fmt.Println("Race over!")
}

func addCarWithChan(name string, c chan<- string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		c <- name
		fmt.Println(name, "added to the race!")
	}
	close(c)
}

func startRaceWithSelect(ferrari <-chan string, lamborghini <-chan string) {
	for {
		select {
		case car, ok := <-ferrari:
			if !ok {
				fmt.Println("Ferrari channel closed!")
				ferrari = nil
				break
			}
			fmt.Println(car, "is racing...")

		case car, ok := <-lamborghini:
			if !ok {
				fmt.Println("Lamborghini channel closed!")
				lamborghini = nil
				break
			}
			fmt.Println(car, "is racing...")
		}

		if ferrari == nil && lamborghini == nil {
			break
		}
		time.Sleep(time.Second)
	}

	fmt.Println("All cars have finished the race!")
}
