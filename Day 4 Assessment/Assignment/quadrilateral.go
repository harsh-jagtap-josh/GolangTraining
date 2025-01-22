package main

import "fmt"

type quadrilateral interface {
	Area() int
	Perimeter() int
}

type rectangleShape struct {
	length  int
	breadth int
}

type squareShape struct {
	side int
}

func (rect rectangleShape) Area() int {
	return rect.length * rect.breadth
}

func (rect rectangleShape) Perimeter() int {
	return 2 * (rect.length + rect.breadth)
}

func (square1 squareShape) Area() int {
	return square1.side * square1.side
}

func (square1 squareShape) Perimeter() int {
	return 4 * square1.side
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
		fmt.Println("Error: ", err)
	} else {

		newRectangle := rectangleShape{
			length:  50,
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
}
