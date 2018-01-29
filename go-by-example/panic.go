package main

import (
	"os"
)

func main() {
	// panic("A problem")

	_, err := os.Create("~/Desktop/file")
	if err != nil {
		panic(err)
	}
}
