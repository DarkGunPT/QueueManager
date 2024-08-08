# QueueManager
In this repo i'm going to be implementing some of the basic functionality of the Queue data structure in Go.
```type Queue struct{}```

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
