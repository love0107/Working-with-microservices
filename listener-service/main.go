package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect rabbit mq
	rabbitConn, err := connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	fmt.Println("Connected to RabbitMQ!")

	// start listening for messages

	// create consumer
	// watch the queue and consume events

}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection
	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")
		time.Sleep(backOff)
		continue
	}
	return connection, nil
}
