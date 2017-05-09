package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// Try to connect
	conn, err := amqp.Dial("amqp://stu:MonkeesUnited@talk.baka-kawaii.de:5672/")
	failOnError(err, "Failed to Connect")
	defer conn.Close()
	// Try to open a Channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open Channel")
	defer ch.Close()
	// Try to declare a queue
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")
	body := "world"
	err = ch.Publish("",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish message")
}
