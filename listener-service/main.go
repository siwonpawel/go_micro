package main

import (
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

	
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
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
