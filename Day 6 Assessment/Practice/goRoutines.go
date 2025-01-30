package main

import (
	"fmt"
)

// Go Routines are basically functions or methods that run concurrently with other functions, hence helping in achieving concurrency.
// extremely cheap, there might be one thread with thousands of Go routines, hence is also called as LIGHT WEIGHT THREADS.
// Routines basically COMMUNICATE USING CHANNELS

// SYNTAX: go <function_name>

// basically, when a routine is called the function goes in the background for execution, and rest of the program keeps working as it was.

func main() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	// time.Sleep(2 * time.Second)

	fmt.Println("Hello World")

}

// in this example, if we run the code... only Hello WOrld is printed, because all the routines in for loop will go behind for execution
// and while they gets executed the main function is completed and exits.. Thus, there also might be output if any of the routines gets finished while the main is executed..
// Thus, sleep of 2 seconds gives time for all the routines to get executed and the output is printed, but with a CATCH
// the sequence of all these routines is RANDOM because GOLANG SCHEDULER manages GO ROUTINES in RANDOM ORDER...

// But we cannot give a time delay everytime we want all the routines to be finished.. Hence we use WAITGROUPS..
// WaitGroup is a "type" of sync package ( type WaitGroup struct {} with only 3 methods in it), is used to wait for the program to finish all the goRoutines launched...
// It works just like COUNTERS, like a while loop.. when a routine is started we increment the counter, and eventually when the counter is 0, the wait is stopped and rest of the program is executed.

// WAIT GORUPS Mainly have only three methods:
// - Add() -> to add a routine in the program, basically increments the counter by 1
// - Done() -> used along with "defer" to just after the Add() to signify that a routine is finished, decrements the counter by 1
// - Wait() -> used at the end of the program constantly checks if the counter is 0, whenever the counter gets 0, calling Wait() returns the
