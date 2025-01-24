package main

import "fmt"

type quadrilateral interface {
	Area() int
	Perimeter() int
}

type rectangle struct {
	length  int
	breadth int
}

type square struct {
	side int
}

func (rect rectangle) Area() int {
	area := rect.length * rect.breadth
	return area
}

func (rect rectangle) Perimeter() int {
	perimeter := 2 * (rect.length + rect.breadth)
	return perimeter
}

func (square1 square) Area() int {
	area := square1.side * square1.side
	return area
}

func (square1 square) Perimeter() int {
	perimeter := 4 * square1.side
	return perimeter
}

func print(newQuadrilateral quadrilateral) {
	fmt.Println("Area: ", newQuadrilateral.Area())
	fmt.Println("Perimeter: ", newQuadrilateral.Perimeter())
}

func main() {

	var userInput int
	fmt.Print("Enter a Number: ")
	_, err := fmt.Scanln(&userInput)

	if err != nil {
		fmt.Errorf("Error in input: ", err)
	}

	newRectangle := rectangle{
		length:  50,
		breadth: 10,
	}

	newSquare := square{
		side: 10,
	}

	if userInput == 1 {
		print(newRectangle)
	} else if userInput == 2 {
		print(newSquare)
	} else {
		fmt.Println("Invalid Input!")
	}
}
