package main

import "fmt"

func main() {

	var days = map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	var day int

	fmt.Print("Enter the Number: ")
	fmt.Scan(&day)

	value, exists := days[day]

	if exists {
		fmt.Println("The day stored at index ", day, " is ", value)
	} else {
		fmt.Println("Not a day!")
	}
	// if day > 8 && day > 0 {
	// 	fmt.Println("Not a day")

	// } else {
	// 	fmt.Println("The day stored at index ", day, " is ", days[day])
	// }
}
