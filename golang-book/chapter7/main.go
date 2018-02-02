package main

import (
	"fmt"
)

func average(xs []float64) float64 {
	//panic("Not implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func add(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	xs := []float64{1, 2, 3, 4, 5}

	fmt.Println(average(xs))
	fmt.Println(add(xs...)) // array... means expanded version of this

	// Closures
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	// Defer, Panic & Recover
	// defer second()
	// first()

	// Panic & Recover
	defer func() {
		str := recover() // used to handle run time panics
		fmt.Println(str)
	}()
	panic("PANIC")

}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}
