package tools

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitServer struct {
	Dial string
	DeclareName string
	Body string
}

var Channel *amqp.Channel

func (rabbit RabbitServer) InitRabbit() *amqp.Channel{
	if Channel!=nil {
		return Channel
	}
	conn, err := amqp.Dial(rabbit.Dial)
	failOnError(err, "Failed to connect to RabbitMQ")
	ch,err := conn.Channel()
	failOnError(err,"Faild To open Channel")
	Channel = ch
	return  Channel
}


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}