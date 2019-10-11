package utils

import (
	"log"

	"github.com/streadway/amqp"
)

// FailOnError reports errors
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// ConnectMQ connects to Rabbit
func ConnectMQ() *amqp.Connection {
	// Establish connection
	conn, err := amqp.Dial("amqp://myuser:mypass@localhost:5672/myvhost")
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}
