package trees

type BTreeNode[T any] struct {
	value T
	left  *BTreeNode[T]
	right *BTreeNode[T]
}
