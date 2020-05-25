package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err,"Failed to open a channel")
	defer ch.Close()

	q,err := ch.QueueDeclare(
		"queue",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err,"Failed To declare a queue")
	//**********************************************
	//传值
	body := bodyFrom(os.Args)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			DeliveryMode:    amqp.Persistent,
			Body:            []byte(body),
		},
		)
	failOnError(err,"Failed To pushlich Qname")
	log.Printf("[x] Sent %s",body)
}
func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "go run this params"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}