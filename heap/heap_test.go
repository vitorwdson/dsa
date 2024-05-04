package heap

import "testing"

func TestHeap(t *testing.T) {
	heap := MinHeap[int]{
		list: []int{},
	}

	if _, err := heap.Pop(); err == nil {
		t.Fatal("Popping from an empty heap should error")
	}

	heap.Insert(5)
	heap.Insert(20)
	heap.Insert(69)
	heap.Insert(42)
	heap.Insert(27)
	heap.Insert(18)
	heap.Insert(100)
	heap.Insert(23)
	heap.Insert(11)

	if v, err := heap.Pop(); err != nil || v != 5 {
		t.Fatalf("Expected return value to be (5, nil) but got (%d, %v)", v, err)
	}
	if v, err := heap.Pop(); err != nil || v != 11 {
		t.Fatalf("Expected return value to be (11, nil) but got (%d, %v)", v, err)
	}
	if v, err := heap.Pop(); err != nil || v != 18 {
		t.Fatalf("Expected return value to be (18, nil) but got (%d, %v)", v, err)
	}
	if v, err := heap.Pop(); err != nil || v != 20 {
		t.Fatalf("Expected return value to be (20, nil) but got (%d, %v)", v, err)
	}
	if v, err := heap.Pop(); err != nil || v != 23 {
		t.Fatalf("Expected return value to be (23, nil) but got (%d, %v)", v, err)
	}
	if v, err := heap.Pop(); err != nil || v != 27 {
		t.Fatalf("Expected return value to be (27, nil) but got (%d, %v)", v, err)
	}
	if v, err := heap.Pop(); err != nil || v != 42 {
		t.Fatalf("Expected return value to be (42, nil) but got (%d, %v)", v, err)
	}
	if v, err := heap.Pop(); err != nil || v != 69 {
		t.Fatalf("Expected return value to be (69, nil) but got (%d, %v)", v, err)
	}
	if v, err := heap.Pop(); err != nil || v != 100 {
		t.Fatalf("Expected return value to be (100, nil) but got (%d, %v)", v, err)
	}
}
