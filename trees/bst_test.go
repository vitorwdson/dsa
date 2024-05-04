package trees

import (
	"slices"
	"testing"
)

func TestBSTree(t *testing.T) {
	tree := &BSTree[int]{}

	tree.Insert(5)
	tree.Insert(20)
	tree.Insert(69)
	tree.Insert(42)
	tree.Insert(27)
	tree.Insert(18)
	tree.Insert(100)
	tree.Insert(23)
	tree.Insert(11)

	expected := []int{5, 11, 18, 20, 23, 27, 42, 69, 100}
	if result := tree.head.WalkOrdered(); !slices.Equal(result, expected) {
		t.Fatalf("Insert - expected %v but got %v", expected, result)
	}

	tree.Delete(18)
	tree.Delete(5)
	tree.Delete(100)
	tree.Delete(69)
	expected = []int{11, 20, 23, 27, 42}
	if result := tree.head.WalkOrdered(); !slices.Equal(result, expected) {
		t.Fatalf("Delete - expected %v but got %v", expected, result)
	}

	if !tree.Find(11) {
		t.Fatal("Find - 11 should be on the tree, but isn't")
	}
	if !tree.Find(20) {
		t.Fatal("Find - 20 should be on the tree, but isn't")
	}
	if !tree.Find(23) {
		t.Fatal("Find - 23 should be on the tree, but isn't")
	}
	if !tree.Find(27) {
		t.Fatal("Find - 27 should be on the tree, but isn't")
	}

	if tree.Find(5) {
		t.Fatal("Find - 5 shouldn't be on the tree, but is")
	}
	if tree.Find(18) {
		t.Fatal("Find - 18 shouldn't be on the tree, but is")
	}
	if tree.Find(100) {
		t.Fatal("Find - 100 shouldn't be on the tree, but is")
	}
	if tree.Find(69) {
		t.Fatal("Find - 69 shouldn't be on the tree, but is")
	}
	if tree.Find(3333) {
		t.Fatal("Find - 3333 shouldn't be on the tree, but is")
	}
}
