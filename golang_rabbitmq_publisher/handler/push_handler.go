package handler

import (
	"encoding/json"

	"net/http"

	constantName "github.com/ekokurniadi/golang_rabbitmq/constant_name"
	"github.com/ekokurniadi/golang_rabbitmq/helper"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type pushHandler struct {
	/// get amqp channel from main
	channel *amqp.Channel
}

func NewPushHandler(channel *amqp.Channel) *pushHandler {
	/// fill channel to pushHandler struct
	return &pushHandler{channel: channel}
}

/// handle request from client
func (h *pushHandler) GetFromRequest(c *gin.Context) {

	/// initialize queue name
	q, err := h.channel.QueueDeclare(
		constantName.Rabbit.QueueName, // name
		true,
		false,
		false,
		false,
		nil,
	)

	/// check if queue is not declared
	if err != nil {
		response := helper.ApiResponse("Get data failed", http.StatusInternalServerError, "error", gin.H{"message": "error when declare queue"})
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	/// get data from request
	var input = c.Query("message")

	/// mapping response to json
	response := helper.ApiResponse("Get data successfully", http.StatusOK, "success", gin.H{"message": input})

	/// convert json to byte
	data, err := json.Marshal(response)

	/// check if json is not converted
	if err != nil {
		response := helper.ApiResponse("Get data failed", http.StatusInternalServerError, "error", gin.H{"message": "decode json failed"})
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	/// mapping message to queue
	message := amqp.Publishing{
		ContentType: constantName.Rabbit.ContentTypeJSON,
		Type:        "get_test",
		Body:        []byte(data),
	}

	/// publish message to queue
	err = h.channel.Publish(
		constantName.Rabbit.ExchangeName,
		q.Name,
		false,
		false,
		message,
	)

	/// check if message is not published
	if err != nil {
		response := helper.ApiResponse("Get data failed", http.StatusInternalServerError, "error", gin.H{"message": "error when publish message"})
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	/// send feed back response to client
	c.JSON(http.StatusOK, response)
}
