# QueueManager
In this repo i'm going to be implementing some of the basic functionality of the Queue data structure in Go.
```
type Consumer struct {
	queue []Flight
}
```
The struct i will use in the queue is the struct Flight:
```
type Flight struct {
  Origin      string
  Destination string
  Price       int
}
```

## Push

This method will take in a Flight and push the flight onto the back of our Items queue.

```func (q *Queue) Push(f Flight)```

## Peek

This method will allow us to view what item is at the front of our queue but not modify the underlying stack values.

```func (q *Queue) Peek() Flight```

## Pop

This method will allow us to pop an element off the front of our Items queue and return to us the first flight.

```func (q *Queue) Pop() Flight```

## IsEmpty

This method will allow us to check if the queue is empty and if so return true

```func (c *Consumer) IsEmpty() bool```

## Tests

We have a tests folder where we can do every test to check if the code is working like he is supposed to work

## Output

```
2024/08/08 03:39:42 Doing a Push
2024/08/08 03:39:42 Doing a Push
2024/08/08 03:39:42 Doing a Push
2024/08/08 03:39:42 Doing a Peek
2024/08/08 03:39:42 The first flight in the queue is: consumers.Flight{Origin:"Porto", Destination:"London", Price:45}
2024/08/08 03:39:42 The queue is: []consumers.Flight{consumers.Flight{Origin:"Porto", Destination:"London", Price:45}, consumers.Flight{Origin:"Lisbon", Destination:"London", Price:85}, consumers.Flight{Origin:"London", Destination:"Porto", Price:55}}
2024/08/08 03:39:42 Doing a Pop
2024/08/08 03:39:42 Removing fly: consumers.Flight{Origin:"Porto", Destination:"London", Price:45} from queue
2024/08/08 03:39:42 The queue is: []consumers.Flight{consumers.Flight{Origin:"Lisbon", Destination:"London", Price:85}, consumers.Flight{Origin:"London", Destination:"Porto", Price:55}}
2024/08/08 03:39:42 Doing a Peek
2024/08/08 03:39:42 The first flight in the queue is: consumers.Flight{Origin:"Lisbon", Destination:"London", Price:85}
2024/08/08 03:39:42 The queue is: []consumers.Flight{consumers.Flight{Origin:"Lisbon", Destination:"London", Price:85}, consumers.Flight{Origin:"London", Destination:"Porto", Price:55}}
2024/08/08 03:39:42 Doing a Pop
2024/08/08 03:39:42 Removing fly: consumers.Flight{Origin:"Lisbon", Destination:"London", Price:85} from queue
2024/08/08 03:39:42 The queue is: []consumers.Flight{consumers.Flight{Origin:"London", Destination:"Porto", Price:55}}
```