package main

import (
	"fmt"
	"math"
)

func calcAreaOfCircle(radius float64) float64 {

	const pi = math.Pi

	var area float64

	area = pi * math.Pow(radius, 2)

	ratio := math.Pow(10, 2)
	return math.Round(float64(area)*ratio) / ratio

	// s := fmt.Sprintf("%.2f", area)

	// return s
}

func main() {
	var radius float64

	fmt.Print("Enter the Radius: ")
	fmt.Scan(&radius)
	area := calcAreaOfCircle(radius)

	fmt.Println("Area of circle is: ", area)

	// fmt.Println("Area of circle is: ", "%.2f", area)
}
