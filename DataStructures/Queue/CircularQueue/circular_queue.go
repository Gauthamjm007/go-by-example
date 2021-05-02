package circular_queue

import "fmt"

type CircularQueue struct {
	data     []interface{}
	head     int
	tail     int
	capacity int
}

//craete a new circular queue
func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}

	return &CircularQueue{
		data:     make([]interface{}, n),
		head:     0,
		tail:     0,
		capacity: n,
	}
}

// IsEmpty returns true if queue is empty
func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

// isFull returns true if queue is full
func (q *CircularQueue) IsFull() bool {
	return q.head == (q.tail+1)%q.capacity
}

func (q *CircularQueue) Enqueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true

}

// Dequeue fetches a element from queue
func (q *CircularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.capacity

	return v
}

// String prints the queue
func (q *CircularQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = q.head
	for {
		result += fmt.Sprintf("<-%+v", q.data[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	return result
}
