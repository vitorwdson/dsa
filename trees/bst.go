package trees

import "cmp"

type BSTree[T cmp.Ordered] struct {
	head *BTreeNode[T]
}

func (t *BSTree[T]) insertAt(value T, ref *BTreeNode[T]) {
	if value > ref.value {
		if ref.right == nil {
			ref.right = &BTreeNode[T]{
				value: value,
			}
			return
		}
		t.insertAt(value, ref.right)
		return
	}

	if ref.left == nil {
		ref.left = &BTreeNode[T]{
			value: value,
		}
		return
	}
	t.insertAt(value, ref.left)
	return
}

func (t *BSTree[T]) Insert(value T) {
	if t.head == nil {
		t.head = &BTreeNode[T]{
			value: value,
		}
		return
	}

	t.insertAt(value, t.head)
}

func (t *BSTree[T]) findNode(value T, ref *BTreeNode[T]) *BTreeNode[T] {
	if ref == nil {
		return nil
	}

	if ref.value == value {
		return ref
	}

	if value > ref.value {
		return t.findNode(value, ref.right)
	}

	return t.findNode(value, ref.left)
}

func (t *BSTree[T]) Find(value T) bool {
	return t.findNode(value, t.head) != nil
}

func (t *BSTree[T]) deleteNode(node *BTreeNode[T], parent *BTreeNode[T], right bool) {
	var newNode *BTreeNode[T]

	if node.left == nil && node.right == nil {
		newNode = nil
	} else if node.left != nil && node.right == nil {
		newNode = node.left
	} else if node.left == nil && node.right != nil {
		newNode = node.right
	} else if node.left != nil && node.right != nil {
		var biggestSmallParent *BTreeNode[T]
		biggestSmall := node.left

		for {
			if biggestSmall.right == nil {
				break
			}
			biggestSmallParent = biggestSmall
			biggestSmall = biggestSmall.right
		}

		t.deleteNode(biggestSmall, biggestSmallParent, true)
		newNode = &BTreeNode[T]{
			value: biggestSmall.value,
			left:  node.left,
			right: node.right,
		}
	}

	if parent == nil {
		t.head = newNode
		return
	}

	if right {
		parent.right = newNode
	} else {
		parent.left = newNode
	}
}

func (t *BSTree[T]) deleteAt(value T, node *BTreeNode[T], parent *BTreeNode[T], right bool) {
	if node == nil {
		return
	}

	if value > node.value {
		t.deleteAt(value, node.right, node, true)
		return
	} else if node.value > value {
		t.deleteAt(value, node.left, node, false)
		return
	}

	t.deleteNode(node, parent, right)
}

func (t *BSTree[T]) Delete(value T) {
	t.deleteAt(value, t.head, nil, false)
}
