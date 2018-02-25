package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (human Human) SayHi() {
	fmt.Printf("Hi, I am %s. You can call me on %s.\n", human.name, human.phone)
}

func main() {
	human := Human{"ratul", 12, "0171"}
	human.SayHi()
}
