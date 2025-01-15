package main

import "fmt"

func main(){

	var days = map[int]string {
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

	if day > 8 {
		fmt.Println("Not a day")

	} else {
		fmt.Println("The day stored at index ", day, " is ", days[day])
	
	}
	// fmt.Println("Hello World")
}
