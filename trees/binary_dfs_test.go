package trees

import (
	"slices"
	"testing"
)

func TestDFSWalk(t *testing.T) {
	h := &BTreeNode[int]{value: 1}
	h.left = &BTreeNode[int]{value: 2}
	h.right = &BTreeNode[int]{value: 3}
	h.left.left = &BTreeNode[int]{value: 4}
	h.left.right = &BTreeNode[int]{value: 5}
	h.right.left = &BTreeNode[int]{value: 6}
	h.right.right = &BTreeNode[int]{value: 7}

	expected := []int{1, 2, 4, 5, 3, 6, 7}
	if returned := h.WalkPre(); !slices.Equal(returned, expected) {
		t.Fatalf("WalkPre expected %v, but got %v", expected, returned)
	}

	expected = []int{4, 5, 2, 6, 7, 3, 1}
	if returned := h.WalkPost(); !slices.Equal(returned, expected) {
		t.Fatalf("WalkPost expected %v, but got %v", expected, returned)
	}

	expected = []int{4, 2, 5, 1, 6, 3, 7}
	if returned := h.WalkOrdered(); !slices.Equal(returned, expected) {
		t.Fatalf("WalkOrdered expected %v, but got %v", expected, returned)
	}
}
