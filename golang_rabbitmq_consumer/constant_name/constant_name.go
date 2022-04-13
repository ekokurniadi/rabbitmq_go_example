package constantName

var queueName string = "QueueService1"
var qmqpServerURL string = "amqp://guest:guest@localhost:5672/"

type rabbitMQ struct {
	QueueName     string
	AmqpServerURL string
}

var RabbitHelper = rabbitMQ{
	QueueName:     queueName,
	AmqpServerURL: qmqpServerURL,
}
