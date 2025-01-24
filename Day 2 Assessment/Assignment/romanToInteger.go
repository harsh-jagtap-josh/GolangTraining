package main

import "fmt"

func convertRomanToInteger(roman string) int {
	twos := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	ones := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	integer := 0
	i := 0
	for i < len(roman)-1 {
		value, exists := twos[roman[i:i+2]]
		if exists {
			integer = integer + value
			i += 2
		} else {
			integer += ones[roman[i:i+1]]
			i++
		}
	}

	if i < len(roman) {
		integer += ones[roman[i:i+1]]
		i++
	}

	return integer
}

func main() {
	var roman string

	fmt.Print("Enter the Roman Number: ")
	fmt.Scan(&roman)

	integer := convertRomanToInteger(roman)

	fmt.Println("The Integer Number for Roman Number:", roman, " is: ", integer)
}
