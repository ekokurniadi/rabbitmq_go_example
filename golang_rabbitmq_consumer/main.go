package main

import (
	"encoding/json"
	constantName "golang_rabbitmq_consume/constant_name"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	amqpServerURL := constantName.RabbitHelper.AmqpServerURL

	rabbitInstance, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer rabbitInstance.Close()

	channel, err := rabbitInstance.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	messages, err := channel.Consume(
		constantName.RabbitHelper.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println(" [*] Waiting for messages. To exit press CTRL+C")

	forever := make(chan bool)

	data := map[string]interface{}{}

	go func() {
		for message := range messages {
			log.Printf(" [*] Received message: %s\n", message.Body)
			log.Printf(" [*] Received type: %s\n", message.Type)
			json.Unmarshal(message.Body, &data)
		}
	}()

	<-forever
}
