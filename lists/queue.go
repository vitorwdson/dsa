package lists

import (
	"errors"
)

type Queue[T any] struct {
	head   *ListNode[T]
	tail   *ListNode[T]
	length int
}

func (q *Queue[T]) Push(v T) {
	node := &ListNode[T]{
		value: v,
	}

	q.length++
	if q.length == 1 {
		q.head = node
		q.tail = node
		return
	}

	q.tail.next = node
	node.prev = q.tail
	q.tail = node
}

func (q *Queue[T]) Pop() (T, error) {
	if q.head == nil {
		return *new(T), errors.New("Cannot pop from empty queue")
	}

	v := q.head.value

	if q.head.next != nil {
		q.head.next.prev = nil
	}
	q.head = q.head.next
	q.length--

	return v, nil
}

func (q Queue[T]) Len() int {
	return q.length
}
