package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	var limit int
	fmt.Println("How far do you want to count Fibonacci to?")
	fmt.Scan(&limit)

	c := make(chan int)
	quit := make(chan int)

	// This starts calculating in the background
	go func() {
		for i := 0; i < limit; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	// This method starts listening to `quit` channel. See the select for-select
	fibonacci(c, quit)
}
