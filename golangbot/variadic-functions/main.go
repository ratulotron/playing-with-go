package main

import "fmt"

func find(number int, numbers ...int) string {
	for index, key := range numbers {
		if key == numbers[index] {
			return fmt.Sprintf("Found %d in %d position.", key, index+1)
		}
	}
	return "Not found"
}

func main() {
	fmt.Println(find(1, 1, 3, 4))
}
