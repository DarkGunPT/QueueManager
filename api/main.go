package main

import (
	"context"
	"log"
	"test/consumers"
)

func main() { // Added two simple tests to see some output; This should only be used to initialize db, config, consumers, publishers, cron jobs, etc...
	ctx := context.Background()

	flightQueueConsumer := consumers.NewConsumer()

	// Test 1

	flight := consumers.Flight{
		Origin:      "Porto",
		Destination: "London",
		Price:       45,
	}

	err := flightQueueConsumer.Consume(ctx, flight, "push")
	if err != nil {
		log.Println(err)
	}

	flight = consumers.Flight{
		Origin:      "Lisbon",
		Destination: "London",
		Price:       85,
	}

	err = flightQueueConsumer.Consume(ctx, flight, "push")
	if err != nil {
		log.Println(err)
	}

	flight = consumers.Flight{
		Origin:      "London",
		Destination: "Porto",
		Price:       55,
	}

	err = flightQueueConsumer.Consume(ctx, flight, "push")
	if err != nil {
		log.Println(err)
	}

	err = flightQueueConsumer.Consume(ctx, flight, "peek")
	if err != nil {
		log.Println(err)
	}

	err = flightQueueConsumer.Consume(ctx, flight, "pop")
	if err != nil {
		log.Println(err)
	}

	err = flightQueueConsumer.Consume(ctx, flight, "peek")
	if err != nil {
		log.Println(err)
	}

	err = flightQueueConsumer.Consume(ctx, flight, "pop")
	if err != nil {
		log.Println(err)
	}

}
