package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// 连接到RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建一个Channel，通过它我们能够调用大多数的API（完成我们的发送、接收任务）
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 定义一个队列
	q, err := ch.QueueDeclare(
		"hello",      // name
		false,     // durable
		false,  // delete when unused
		false,   // exclusive
		false,     // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 发送消息
	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
