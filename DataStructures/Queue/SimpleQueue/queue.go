package queue

import "fmt"

type Queue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

//create a new Queue
func NewQueue(n int) *Queue {
	return &Queue{
		data:     make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// Add a new queue at the end of the queue
func (q *Queue) Enqueue(v interface{}) bool {
	if q.tail == q.capacity {
		if q.head == 0 {
			return false
		}
		// move data
		for i := q.head; i < q.tail; i++ {
			q.data[i-q.head] = q.data[i]
		}
		q.tail -= q.head
		q.head = 0
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

// Dequeue fetches a element from queue
func (q *Queue) Dequeue() interface{} {
	if q.head == q.tail {
		return nil
	}
	val := q.data[q.head]
	q.head++
	return val
}

// String prints the queue
func (q *Queue) String() string {
	if q.head == q.tail {
		return "empty queue"
	}
	result := "head"
	for i := q.head; i <= q.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", q.data[i])
	}
	result += "<-tail"
	return result
}
