package main

import (
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	rabbitConn, err := connect()
	if err != nil {
		log.Println("err")
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ")

	log.Println("Listening for and consuming RabbitMQ messages...")
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			log.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			return c, nil
		}

		if counts > 5 {
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2))
		log.Println("backing off...")
		time.Sleep(backOff)
	}
}
