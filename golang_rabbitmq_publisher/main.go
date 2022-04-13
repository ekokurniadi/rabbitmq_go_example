package main

import (
	constantName "github.com/ekokurniadi/golang_rabbitmq/constant_name"
	"github.com/ekokurniadi/golang_rabbitmq/handler"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {

	/// connect to RabbitMQ
	amqpServerURL := constantName.Rabbit.AmqpServerURL

	/// create rabbitMQ connection
	rabbitInstance, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer rabbitInstance.Close()

	/// create rabbitMQ channel
	channel, err := rabbitInstance.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	/// initialize router using gin gonic
	router := gin.Default()

	/// initialize push handler and send channel to handler
	pushHandler := handler.NewPushHandler(channel)

	/// handle request from client
	router.GET("/test", pushHandler.GetFromRequest)

	/// run server
	router.Run()
}
