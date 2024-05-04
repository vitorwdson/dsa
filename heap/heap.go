package heap

import (
	"cmp"
	"errors"
)

type MinHeap[T cmp.Ordered] struct {
	list []T
}

func (h *MinHeap[T]) bubbleUp(idx int) {
	if idx == 0 {
		return
	}

	upperIdx := (idx - 1) / 2
	currValue := h.list[idx]

	if h.list[upperIdx] < currValue {
		return
	}

	h.list[idx] = h.list[upperIdx]
	h.list[upperIdx] = currValue
	h.bubbleUp(upperIdx)
}

func (h *MinHeap[T]) Insert(value T) {
	h.list = append(h.list, value)
	h.bubbleUp(len(h.list) - 1)
}

func (h *MinHeap[T]) bubbleDown(idx int) {
	leftIdx := 2*idx + 1
	rightIdx := 2*idx + 2
	length := len(h.list)

	var smallestValue T
	var smallestIdx int

	if leftIdx < length && rightIdx < length {
		leftValue := h.list[leftIdx]
		rightValue := h.list[rightIdx]

		if leftValue < rightValue {
			smallestIdx = leftIdx
			smallestValue = leftValue
		} else {
			smallestIdx = rightIdx
			smallestValue = rightValue
		}
	} else if leftIdx < length && rightIdx >= length {
		smallestIdx = leftIdx
		smallestValue = h.list[leftIdx]
	} else if leftIdx >= length && rightIdx >= length {
		return
	}

	currValue := h.list[idx]
	if smallestValue >= currValue {
		return
	}

	h.list[idx] = smallestValue
	h.list[smallestIdx] = currValue
	h.bubbleDown(smallestIdx)
}

func (h *MinHeap[T]) Pop() (T, error) {
	if len(h.list) == 0 {
		return *new(T), errors.New("Cannot pop from empty heap")
	}

	poppedValue := h.list[0]
	h.list[0] = h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]
	h.bubbleDown(0)

	return poppedValue, nil
}
