package publisher

import (
	"log"
	"os"
	"strings"

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
		"task_queue",      // name
		true,   // 持久化
		false,  // delete when unused
		false,   // exclusive
		false,     // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 发送消息
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",           // 使用默认的Exchange
		q.Name,       // routing key
		false,        // mandatory
		false,
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,   // 消息持久化
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
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
