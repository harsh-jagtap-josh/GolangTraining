// In Go (also known as Golang), the concepts of panic and recover are used to handle exceptional situations, somewhat similar to exceptions in other programming languages. However, Go's approach to handling errors is different from traditional try-catch exception handling found in languages like Java or Python. Let's take a look at how panic and recover work internally in Go.
// Definition: In Go programming, a panic signifies a critical runtime error that disrupts the normal execution flow. Distinct from conventional error handling, panics are used for exceptional errors that the usual program logic cannot easily anticipate or handle, such as nil pointer dereferences or out-of-range array access.
// Panic:

// A panic occurs when a Go program encounters a runtime error that it cannot recover from. This could be due to various reasons, such as division by zero, attempting to access an out-of-bounds array index, or encountering a nil pointer dereference.
// When a panic occurs, the normal flow of the program is stopped immediately, and the program starts unwinding the stack, which means it starts to undo the function calls, similar to how a stack unwinds during a function call chain.
// While unwinding the stack, deferred function calls are executed in the reverse order of their registration (LIFO order).
// If no deferred function handles the panic, the program terminates, and an error message is printed to the standard error stream.

// Recover:

// To prevent a panic from crashing the entire program, Go provides the recover() function.
// The recover() function is used in combination with the defer statement. When recover() is called inside a deferred function, it stops the panic's propagation and returns the value that was passed to the call to panic().
// If recover() is called outside the deferred function or if it's not called at all, it will return nil.
// To use recover() effectively, you need to defer a function that includes a call to recover() at the beginning, and this function should be placed in the appropriate location in the call stack to capture and handle the panic.

package main

import "fmt"

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}
func someFunction() {
	defer recoverFromPanic() // Using defer to call the recovery function
	// Simulate a panic
	panic("Something went wrong!")
}

func main() {
	fmt.Println("Start")
	someFunction()
	fmt.Println("End")
}
