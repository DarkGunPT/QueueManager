package consumers

import (
	"context"
	"errors"
	"log"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type Consumer struct {
	queue []Flight
}

func NewConsumer() *Consumer {
	return &Consumer{
		queue: make([]Flight, 0),
	}
}

func (c *Consumer) Consume(ctx context.Context, flight Flight, action string) error { // For this challenge purpose i will be using one string for to identify the action needed in the queue
	switch action {
	case "push":
		{
			log.Println("Doing a Push")
			c.Push(flight)
		}
	case "peek":
		{
			log.Println("Doing a Peek")
			flight, err := c.Peek()
			if err != nil {
				return err
			}
			log.Printf("The first flight in the queue is: %#v", flight)
			log.Printf("The queue is: %#v", c.queue)
		}
	case "pop":
		{
			log.Println("Doing a Pop")
			flight, err := c.Pop()
			if err != nil {
				return err
			}
			log.Printf("Removing fly: %#v from queue", flight)
			log.Printf("The queue is: %#v", c.queue)
		}
	}
	return nil
}

func (c *Consumer) Pop() (Flight, error) {
	if c.IsEmpty() {
		return Flight{}, errors.New("queue is empty")
	}
	firstFlight := c.queue[0]
	c.queue = c.queue[1:]
	return firstFlight, nil
}

func (c *Consumer) Push(f Flight) {
	c.queue = append(c.queue, f)
}

func (c *Consumer) Peek() (Flight, error) {
	if c.IsEmpty() {
		return Flight{}, errors.New("queue is empty")
	}
	return c.queue[0], nil
}

func (c *Consumer) IsEmpty() bool {
	if len(c.queue) == 0 {
		return true
	}
	return false
}
