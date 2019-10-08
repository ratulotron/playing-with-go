package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		// fmt.Println(s)
		fmt.Println(time.Now(), "--", s)
	}
}

func main() {
	go say("world")
	say("hello")
	fmt.Println("finished", time.Now())
}
