package main

import (
	"fmt"
)

type rectangle struct {
	width, height float64
}

//A method is a function with an implicit first argument, called a receiver.
func area(r rectangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := rectangle{12, 2}
	r2 := rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))
}
