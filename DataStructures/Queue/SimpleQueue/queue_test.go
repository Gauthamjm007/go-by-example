package queue

import "testing"

func TestQueue_NewQueue(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)

}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)
	q.Dequeue()
	t.Log(q)
	q.Dequeue()
	t.Log(q)

	t.Log(q.String())
}
