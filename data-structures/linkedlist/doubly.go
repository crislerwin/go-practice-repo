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

func (ll *Doubly[T]) AddAtEnd(val T) {
	ll.lazyInit()
	ll.insertValue(val, ll.Head.Prev)
}

func (ll *Doubly[T]) Remove(n *Node[T]) T {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	n.Next = nil
	n.Prev = nil
	return n.Val
}

func (ll *Doubly[T]) DelAtBeg() (T, bool) {
	if ll.Head.Next == nil {
		var r T
		return r, false
	}
	n := ll.Head.Next
	val := n.Val
	ll.Remove(n)
	return val, true
}

func (ll *Doubly[T]) DelAtEnd() (T, bool) {
	if ll.Head.Prev == nil {
		var r T
		return r, false
	}
	n := ll.Head.Prev
	val := n.Val
	ll.Remove(n)
	return val, true
}
