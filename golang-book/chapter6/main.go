package main

import "fmt"

func main() {
	// var x [5]int
	// x[4] = 100
	// fmt.Println(x)

	// slice1 := []int{1, 2, 3}
	// slice2 := append(slice1, 4, 5)
	// fmt.Println(slice1, slice2)

	// slice1 = []int{1, 2, 3}
	// slice2 = make([]int, 2)
	// copy(slice2, slice1)
	// fmt.Println(slice1, slice2)

	// var m map[string]int
	// ^ only this won't work coz maps gotta be initialized during deplaration
	// m := make(map[string]int)
	// m["key"] = 10
	// m["key2"] = 15
	// fmt.Println(m)

	// elements := map[string]string{
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	// fmt.Println(elements["Li"])

	// This would give nothing coz Un doesn't exist:
	// fmt.Println("Un"])

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	} else {
		fmt.Println("Naikka pai nai")
	}

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := x[0]
	for _, number := range x {
		if number < min {
			min = number
		}
	}
	fmt.Println(min)

}
