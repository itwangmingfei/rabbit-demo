package main

import (

	"github.com/streadway/amqp"
	"log"
	"os"
	"rebbitMQ_demo/tools"
)

func main() {
	body := os.Args
	Rabbitput(body[1])

}
func Rabbitput(body string) {
	var Rabbit tools.RabbitServer
	Rabbit.Dial = "amqp://guest:guest@localhost:5672/"
	Rabbit.DeclareName = "queueName"
	Rabbit.Body = body
	channel := Rabbit.InitRabbit()
	//*******************************************
	q,err := channel.QueueDeclare(Rabbit.DeclareName,false,false,false,false,nil)
	if err!=nil{
		failOnError(err,"Faild To declare  Queue")
	}
	err = channel.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "texp/plain",
		Body:        []byte(Rabbit.Body),
	})
}


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
