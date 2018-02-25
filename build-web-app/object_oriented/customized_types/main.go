package main

import "fmt"

type age int
type money float64
type months map[string]int

func main() {
	m := months{
		"January":  31,
		"February": 28,
		"March":    31,
		"April":    30,
		"May":      31,
		"June":     30,
		"July":     31,
		"August":   30,
	}

	fmt.Println(m)
}
