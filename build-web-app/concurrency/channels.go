package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	a := []int{6, 32, 1, 2, 6, 23, -1}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

// func v(c chan int) {
// 	s1 := rand.NewSource(time.Now().UnixNano())
// 	r1 := rand.New(s1)
// 	c <- r1.Int() // send v to channel ch
// }

// func main() {
// 	ci := make(chan int)
// 	// cs := make(chan string)
// 	// cf := make(chan interface{})

// 	go v(ci)

// 	time.Sleep(2)

// 	n := <-ci
// 	fmt.Println(n)
// }
