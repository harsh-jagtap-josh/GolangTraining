// Copied from Go By Examples :)

// Interfaces are basically behavioral type, it just says that any struct that follows the behavior(methods) defined in the interface is a part of the interface.
// --> the type of the interface can be defined using: interfaceObject(type) or if we use interfaceObject(structName) for any struct that has similar behavior, it returns true or false.
// --> a.(type)
// -->
package main

import (
	"fmt"
	"math"
)

// an interface with some behavior or methods
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Functions defined on interfaces can be used on any of the structs that follow the interface behavior or have all the methods that are in the interface..
// in this function if we add any struct like rectangle, square or circle, it checks whether the passed input structure is a circle or not...
// for checking the type interfaces simply use interfaceObject.(type) to check type, or interfaceObject.(structName) to check if passed structObject is of type structName
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
