package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

//area method for circle
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

//area method
func (r rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	c1 := circle{10}
	c2 := circle{25}
	r1 := rectangle{9, 4}
	r2 := rectangle{12, 2}

	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
}
