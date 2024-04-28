package trees

import "testing"

func TestBTreeCompare(t *testing.T) {
	a := &BTreeNode[int]{value: 1}
	a.left = &BTreeNode[int]{value: 2}
	a.right = &BTreeNode[int]{value: 3}
	a.left.left = &BTreeNode[int]{value: 4}
	a.left.right = &BTreeNode[int]{value: 5}
	a.right.left = &BTreeNode[int]{value: 6}
	a.right.right = &BTreeNode[int]{value: 7}

	b := &BTreeNode[int]{value: 1}
	b.left = &BTreeNode[int]{value: 2}
	b.right = &BTreeNode[int]{value: 4}
	b.left.left = &BTreeNode[int]{value: 4}
	b.left.right = &BTreeNode[int]{value: 42}
	b.right.left = &BTreeNode[int]{value: 6}
	b.right.left.right = &BTreeNode[int]{value: 7}

	c := &BTreeNode[int]{value: 1}
	c.left = &BTreeNode[int]{value: 2}
	c.right = &BTreeNode[int]{value: 3}
	c.left.left = &BTreeNode[int]{value: 4}
	c.left.right = &BTreeNode[int]{value: 5}
	c.right.left = &BTreeNode[int]{value: 6}
	c.right.right = &BTreeNode[int]{value: 7}

	if a.IsEqual(b) {
		t.Fatalf("A should not be equal to B")
	}
	if !a.IsEqual(c) {
		t.Fatalf("A should be equal to C")
	}
}
