package main

import (
	"os"

	service "github.com/ratulotron/gogo/service"
)

func main() {
	port := os.Getenv("POST")
	if len(port) == 0 {
		port = "3000"
	}

	server := service.NewServer()
	server.Run(":" + port)
}
