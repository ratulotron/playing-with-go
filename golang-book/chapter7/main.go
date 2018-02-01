package main

import "fmt"

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
}
