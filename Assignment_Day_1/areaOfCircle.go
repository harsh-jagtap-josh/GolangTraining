package main

import (
	"fmt"
	"math"
)

func calcAreaOfCircle(radius float64) float64 {

	const Pi = math.Pi

	var area float64

	area = Pi * math.Pow(radius, 2)

	ratio := math.Pow(10, 2)
	return math.Round(float64(area)*ratio) / ratio

}

func main() {
	var radius float64

	fmt.Scan(&radius)

	area := calcAreaOfCircle(radius)

	fmt.Print("Area of circle is: ", area)
}
