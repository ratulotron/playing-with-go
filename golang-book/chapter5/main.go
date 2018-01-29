package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1]) // Will print the char code
	// Char is represented as a byte and a byte is an integer
	fmt.Println("Hello " + "World")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
