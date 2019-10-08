package main

import (
	"fmt"
	"strconv"
)

type element interface{}
type list []element
type person struct {
	name string
	age  int
}

func (p person) String() string {
	return "(name : " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(list, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = person{"Dennis", 70}

	// for index, element := range list {
	// 	if value, ok := element.(int); ok {
	// 		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	// 	} else if value, ok := element.(string); ok {
	// 		fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
	// 	} else if value, ok := element.(person); ok {
	// 		fmt.Printf("list[%d] is a person and its value is %s\n", index, value)
	// 	} else {
	// 		fmt.Printf("list[%d] is of a different type\n", index)
	// 	}
	// }
	for index, elem := range list {
		switch value := elem.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its vallue is %s\n", index, value)
		case person:
			fmt.Printf("list[%d] is a person and its vallue is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}
