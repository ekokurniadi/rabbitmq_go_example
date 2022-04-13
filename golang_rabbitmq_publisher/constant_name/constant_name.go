package constantName

var exchangeName string = ""
var queueName string = "QueueService1"
var qmqpServerURL string = "amqp://guest:guest@localhost:5672/"
var contentTypeJSON string = "application/json"
var contentTypeText string = "text/plain"

type rabbitMQ struct {
	ExchangeName    string
	QueueName       string
	AmqpServerURL   string
	ContentTypeJSON string
	ContentTypeText string
}

var Rabbit = rabbitMQ{
	ExchangeName:    exchangeName,
	QueueName:       queueName,
	AmqpServerURL:   qmqpServerURL,
	ContentTypeJSON: contentTypeJSON,
	ContentTypeText: contentTypeText,
}
