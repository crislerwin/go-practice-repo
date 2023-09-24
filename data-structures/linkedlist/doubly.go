package linkedlist

type Doubly[T any] struct {
	Head *Node[T]
}

func (ll *Doubly[T]) Init() *Doubly[T] {
	ll.Head = &Node[T]{}
	ll.Head.Next = ll.Head
	ll.Head.Prev = ll.Head
	return ll
}

func NewDoubly[T any]() *Doubly[T] {
	return new(Doubly[T]).Init()
}

func (ll *Doubly[T]) AddAtBeg(val T) {
	ll.lazyInit()
	ll.insertValue(val, ll.Head)
}

func (ll *Doubly[T]) lazyInit() {
	if ll.Head.Next == nil {
		ll.Init()
	}
}

func (ll *Doubly[T]) insert(n, at *Node[T]) *Node[T] {
	n.Prev = at
	n.Next = at.Next
	n.Prev.Next = n
	n.Next.Prev = n
	return n
}

func (ll *Doubly[T]) insertValue(val T, at *Node[T]) *Node[T] {
	return ll.insert(NewNode(val), at)
}
