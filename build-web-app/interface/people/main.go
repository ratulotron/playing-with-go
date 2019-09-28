package main

import (
	"fmt"
	"strconv"
)

// Human is a human
type Human struct {
	name  string
	age   int
	phone string
}

// Student has Human attributes
type Student struct {
	Human
	school string
	loan   float32
}

// Employee is a human
type Employee struct {
	Human
	company string
	money   float32
}

// Men is a interface for humans
type Men interface {
	sayHi()
	sing(lyrics string)
	guzzle(beerStein string)
}

type youngChap interface {
	sayHi()
	sing(song string)
	BorrowMoney(amount float32)
}

type elderlyGent interface {
	sayHi()
	sing(song string)
	spendSalary(amount float32)
}

func (h Human) sayHi() {
	fmt.Printf("Hi, I am %s. You can call me on %s.\n", h.name, h.phone)
}

func (h Human) sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

func (h Human) guzzle(beerStein string) {
	fmt.Println("Guzzle guzzle guzzle...", beerStein)
}

func (h Human) String() string {
	return "Name: " + h.name + "\n" +
		"Age: " + strconv.Itoa(h.age) + "years\n" +
		"Contact: " + h.phone + "\n"
}

// SayHi is overloaded by Employee
func (e Employee) sayHi() {
	fmt.Printf("Hi, I am %s. I work at %s. Call me on %s.\n", e.name, e.company, e.phone)
}

func (s Student) borrowMoney(amount float32) {
	s.loan += amount
}

func (e Employee) spendSalary(amount float32) {
	e.money -= amount
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-xxx"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	// Define interface i
	var i Men

	// i can store Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.sayHi()
	i.sing("November rain")

	// i can store Employee
	i = tom
	fmt.Println("This is Tom, an Employee: ")
	i.sayHi()
	i.sing("Born to be wild!")

	// Slice of Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.sayHi()
	}

	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
