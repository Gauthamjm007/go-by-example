package circular_queue

import "testing"

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	t.Log(q)
}

func TestCircularDeQqueue(t *testing.T) {
	q1 := NewCircularQueue(5)
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)
	q1.Enqueue(4)
	q1.Enqueue(5)
	t.Log(q1)
	t.Log(q1.Dequeue())
	t.Log(q1)
}
