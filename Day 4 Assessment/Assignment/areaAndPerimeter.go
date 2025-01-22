package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
}

func (rect rectangle) area() int {
	area := rect.length * rect.breadth
	return area
}

func (rect rectangle) perimeter() int {
	perimeter := 2 * (rect.length + rect.breadth)
	return perimeter
}

func main() {

	var length, breadth int
	fmt.Print("Enter the Length: ")
	_, err := fmt.Scanln(&length)
	if err != nil {
		fmt.Errorf("Error: ", err)
	}

	fmt.Print("Enter the Breadth: ")
	_, err2 := fmt.Scanln(&breadth)

	if err2 != nil {
		fmt.Errorf("Error: ", err2)
	}

	newRectangle := rectangle{
		length:  length,
		breadth: breadth,
	}

	fmt.Println("Area of Rectangle is: ", newRectangle.area())
	fmt.Println("Perimeter of Rectangle is: ", newRectangle.perimeter())
}
