package tests

import (
	"context"
	"test/consumers"
	"testing"
)

func TestConsume(t *testing.T) {
	ctx := context.Background()

	flightQueueConsumer := consumers.NewConsumer()

	// Tests

	flight := consumers.Flight{
		Origin:      "Porto",
		Destination: "London",
		Price:       45,
	}

	err := flightQueueConsumer.Consume(ctx, flight, "push")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	flight = consumers.Flight{
		Origin:      "Lisbon",
		Destination: "London",
		Price:       85,
	}

	err = flightQueueConsumer.Consume(ctx, flight, "push")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	flight = consumers.Flight{
		Origin:      "London",
		Destination: "Porto",
		Price:       55,
	}

	err = flightQueueConsumer.Consume(ctx, flight, "push")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	err = flightQueueConsumer.Consume(ctx, flight, "peek")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	err = flightQueueConsumer.Consume(ctx, flight, "pop")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	err = flightQueueConsumer.Consume(ctx, flight, "peek")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	err = flightQueueConsumer.Consume(ctx, flight, "pop")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
