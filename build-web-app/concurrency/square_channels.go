package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}
	close(c)
}

func main() {
	fmt.Println("main() started")
	c := make(chan int)

	go squares(c) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	// for {
	// 	val, ok := <-c
	// 	if ok == false {
	// 		fmt.Println(val, ok, "<--- the loop broke!")
	// 		break // exit break loop
	// 	} else {
	// 		fmt.Println(val, ok)
	// 	}
	// }
	for val := range c {
		fmt.Println(val)
	}
	fmt.Println("main() stopped")
}
