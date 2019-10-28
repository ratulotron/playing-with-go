package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var messageBuffer chan string

func listenAndPrint(c chan string) {
	for res := range c {
		fmt.Println(res)
	}
}

// SendMessage e.GET("/send", SendMessage)
func sendMessage(c echo.Context) error {
	message := c.QueryParam("message")
	messageBuffer <- message
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Received message",
	})
}

func main() {
	messageBuffer = make(chan string)

	go listenAndPrint(messageBuffer)
	// Initialize app
	e := echo.New()
	e.GET("/", sendMessage)
	go e.Logger.Fatal(e.Start(":1323"))
}
