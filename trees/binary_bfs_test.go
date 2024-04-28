package trees

import (
	"slices"
	"testing"
)

func TestBFSWalk(t *testing.T) {
	h := &BTreeNode[int]{value: 1}
	h.left = &BTreeNode[int]{value: 2}
	h.right = &BTreeNode[int]{value: 3}
	h.left.left = &BTreeNode[int]{value: 4}
	h.left.right = &BTreeNode[int]{value: 5}
	h.right.left = &BTreeNode[int]{value: 6}
	h.right.right = &BTreeNode[int]{value: 7}

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	if returned := h.WalkBFS(); !slices.Equal(returned, expected) {
		t.Fatalf("WalkBFS expected %v, but got %v", expected, returned)
	}
}
