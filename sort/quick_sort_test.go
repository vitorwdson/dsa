package sort

import (
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	sample := []int{
		8,
		3,
		10,
		2,
		5,
		1,
	}

	sorted := []int{
		1,
		2,
		3,
		5,
		8,
		10,
	}

	QuickSort(sample)

	if !slices.Equal(sample, sorted) {
		t.Fatalf("Expected %v, but got %v", sorted, sample)
	}
}
