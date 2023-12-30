package main

import (
	"errors"
	"fmt"
)

const MAX_SIZE = 100 // Maximum size of the queue

type Queue struct {
	A          []any
	head, tail int
}

func NewQueue(s int) *Queue {
	q := &Queue{
		head: -1,
		tail: -1,
		A:    make([]any, s),
	}
	return q
}

func (q *Queue) isEmpty() bool {
	return q.head == -1 && q.tail == -1
}

func (q *Queue) isFull() bool {
	if (q.tail+1)%len(q.A) == q.head {
		return true
	} else {
		return false
	}
}

func (q *Queue) Enqueue(e any) error {
	if q.isFull() {
		return errors.New("queue is full")
	}
	if q.isEmpty() {
		q.head, q.tail = 0, 0
	} else {
		q.tail = (q.tail + 1) % len(q.A)
	}
	q.A[q.tail] = e
	return nil
}

func (q *Queue) Dequeue() error {
	if q.isEmpty() {
		return errors.New("queue is empty")
	} else if q.head == q.tail {
		q.head, q.tail = -1, -1
	} else {
		q.head = (q.head + 1) % len(q.A)
	}
	return nil
}

func (q *Queue) Head() (any, error) {
	if q.head == -1 {
		return nil, errors.New("queue is empty")
	}
	return q.A[q.head], nil
}

func main() {
	q := NewQueue(4)
	q.Enqueue("Kudik")
	fmt.Println(q.Head())
}
