package lists

import "errors"

type Queue[T any] struct {
	head   *ListNode[T]
	length int
}

func (q *Queue[T]) Push(v T) {
	node := &ListNode[T]{
		value: v,
	}

	if q.head == nil {
		q.head = node
	} else {
		q.head.next = node
		node.prev = q.head
	}
	q.length++
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
