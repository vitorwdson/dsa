package trees

type BTreeNode[T any] struct {
	value T
	left  *BTreeNode[T]
	right *BTreeNode[T]
}

func (n *BTreeNode[T]) recurseWalkPre(visited []T) []T {
	visited = append(visited, n.value)
	if n.left != nil {
		visited = n.left.recurseWalkPre(visited)
	}
	if n.right != nil {
		visited = n.right.recurseWalkPre(visited)
	}

	return visited
}

func (n *BTreeNode[T]) WalkPre() []T {
	return n.recurseWalkPre([]T{})
}

func (n *BTreeNode[T]) recurseWalkPost(visited []T) []T {
	if n.left != nil {
		visited = n.left.recurseWalkPost(visited)
	}
	if n.right != nil {
		visited = n.right.recurseWalkPost(visited)
	}
	visited = append(visited, n.value)

	return visited
}

func (n *BTreeNode[T]) WalkPost() []T {
	return n.recurseWalkPost([]T{})
}

func (n *BTreeNode[T]) recurseWalkOrdered(visited []T) []T {
	if n.left != nil {
		visited = n.left.recurseWalkOrdered(visited)
	}
	visited = append(visited, n.value)
	if n.right != nil {
		visited = n.right.recurseWalkOrdered(visited)
	}

	return visited
}

func (n *BTreeNode[T]) WalkOrdered() []T {
	return n.recurseWalkOrdered([]T{})
}
