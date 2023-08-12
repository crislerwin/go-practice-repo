package linkedlist

type Singly[T any] struct {
	length int

	Head *Node[T]
}

func NewSingly[T any]() *Singly[T] {
	return &Singly[T]{}
}

func (ll *Singly[T]) AddAtBeg(val T) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}
