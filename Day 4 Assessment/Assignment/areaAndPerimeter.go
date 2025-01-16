package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
}

func (rect rectangle) area() int {
	return rect.length * rect.breadth
}

func (rect rectangle) perimeter() int {
	return 2 * (rect.length + rect.breadth)
}

func main() {

	var length, breadth int
	fmt.Print("Enter the Length: ")
	fmt.Scanln(&length)
	fmt.Print("Enter the Breadth: ")
	fmt.Scanln(&breadth)

	newRectangle := rectangle{
		length:  length,
		breadth: breadth,
	}

	fmt.Println("Area of Rectangle is: ", newRectangle.area())
	fmt.Println("Perimeter of Rectangle is: ", newRectangle.perimeter())
}
