package core

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ is a struct that defines the RabbitMQ connection
type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

// NewRabbitMQ is a function that returns a new RabbitMQ connection
func NewRabbitMQ(scheme string, host string, port int, user string, password string) *RabbitMQ {

	amqpUri := fmt.Sprintf("%s://%s:%s@%s:%d/",
		scheme,
		user,
		password,
		host,
		port,
	)

	_, err := amqp.ParseURI(amqpUri)

	if err != nil {
		panic(fmt.Sprintf("failed to parse the AMQP URI: %s", err))
	}

	// Connect to RabbitMQ
	conn, err := amqp.Dial(amqpUri)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to RabbitMQ: %s", err))
	}

	// Create a channel
	ch, err := conn.Channel()

	if err != nil {
		panic(fmt.Sprintf("failed to open a channel: %s", err))
	}

	return &RabbitMQ{
		Conn: conn,
		Ch:   ch,
	}
}
