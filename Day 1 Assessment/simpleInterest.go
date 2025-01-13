package main

import (
	"fmt"
	"math"
)

func calcSimpInterest(principle float64, rate float64, time float64) float64 {
	var simpleInterest float64 = (principle * rate * time) / 100

	ratio := math.Pow(10, 2)
	return math.Round(float64(simpleInterest)*ratio) / ratio
}

func main() {

	var principle, rate, time float64

	fmt.Print("Enter Principle, Rate, and Time (in years): ")

	fmt.Scanf("%v,%v,%v", &principle, &rate, &time)

	fmt.Println("Principle: ", principle, ", Rate: ", rate, ", Time: ", time)

	simpleInterest := calcSimpInterest(principle, rate, time)

	fmt.Println("Simple Interst is: ", float64(simpleInterest))
}
