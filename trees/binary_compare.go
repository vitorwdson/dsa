package trees

func compareBTrees[T comparable](a, b *BTreeNode[T]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return compareBTrees(a.left, b.left) && compareBTrees(a.right, b.right)
}

func (n *BTreeNode[T]) IsEqual(o *BTreeNode[T]) bool {
	return compareBTrees(n, o)
}
