package trees

import (
	"github.com/vitorwdson/dsa/lists"
)

func (n *BTreeNode[T]) WalkBFS() []T {
	q := lists.Queue[*BTreeNode[T]]{}
	q.Push(n)

	seen := make([]T, 0)
	for q.Len() > 0 {
		node, err := q.Pop()
		if err != nil || node == nil {
			return seen
		}
		seen = append(seen, node.value)
		q.Push(node.left)
		q.Push(node.right)
	}

	return seen
}
