package trees

import (
	"github.com/vitorwdson/dsa/lists"
)

func recurseWalkBFS[T any](q lists.Queue[*BTreeNode[T]], seen []T) []T {
	node, err := q.Pop()
	if err != nil || node == nil {
		return seen
	}

	seen = append(seen, node.value)
	q.Push(node.left)
	q.Push(node.right)

	return recurseWalkBFS(q, seen)
}

func (n *BTreeNode[T]) WalkBFS() []T {
	q := lists.Queue[*BTreeNode[T]]{}
	q.Push(n)

	return recurseWalkBFS(q, make([]T, 0))
}
