package main

import (
	"fmt"

	"github.com/efaraz27/go-auth/email-service/core"
)

func main() {
	// Load the environment variables
	config := core.LoadConfig()

	// Connect to RabbitMQ
	rabbitmq := core.NewRabbitMQ(
		config.AmqpScheme,
		config.AmqpHost,
		config.AmqpPort,
		config.AmqpUser,
		config.AmqpPassword,
	)

	// Create a queue
	queue, err := core.DeclareQueue(rabbitmq.Ch, config.AmqpQueue)
	if err != nil {
		panic("failed to declare a queue")
	}

	// Consume the queue
	msgs, err := rabbitmq.Ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		panic("failed to register a consumer")
	}

	// Listen for messages
	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Printf("Received a message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")

	<-forever

}
