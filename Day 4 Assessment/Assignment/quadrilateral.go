package main

import "fmt"

type quadrilateral interface {
	areaOfShape() int
	perimeterOfShape() int
}

type rectangleShape struct {
	length  int
	breadth int
}

type squareShape struct {
	side int
}

func (rect rectangleShape) areaOfShape() int {
	return rect.length * rect.breadth
}

func (rect rectangleShape) perimeterOfShape() int {
	return 2 * (rect.length + rect.breadth)
}

func (square1 squareShape) areaOfShape() int {
	return square1.side * square1.side
}

func (square1 squareShape) perimeterOfShape() int {
	return 4 * square1.side
}

func print(newQuadrilateral quadrilateral) {
	fmt.Println("Area: ", newQuadrilateral.areaOfShape())
	fmt.Println("Perimeter: ", newQuadrilateral.perimeterOfShape())
}

func main() {

	var userInput int
	fmt.Print("Enter a Number: ")
	fmt.Scanln(&userInput)

	newRectangle := rectangleShape{
		length:  10,
		breadth: 10,
	}

	newSquare := squareShape{
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
