package trees

type BTreeNode[T comparable] struct {
	value T
	left  *BTreeNode[T]
	right *BTreeNode[T]
}
