package lists

type ListNode[T any] struct {
	value T
	next  *ListNode[T]
	prev  *ListNode[T]
}
