package main

import (
	"example/baseProject/messageBroker"
	"fmt"
)

func main() {
	broker, err := messageBroker.NewRabbitMQBroker("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer broker.Close()
	// Consume messages
	queueName := "productlist"
	err = broker.ConsumeMessages(queueName, func(message string) {
		fmt.Println("Received a message:", message)
	})
	if err != nil {
		panic(err)
	}
	select {}
}
