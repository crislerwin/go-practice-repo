package stack

type Node struct {
	Val  any
	Next *Node
}

type Stack struct {
	top    *Node
	length int
}

func (ll *Stack) push(n any) {
	newStack := &Node{}
	newStack.Val = n
	newStack.Next = ll.top
	ll.top = newStack
	ll.length++
}

func (ll *Stack) show() (in []any) {
	current := ll.top
	for current != nil {
		in = append(in, current.Val)
		current = current.Next
	}
	return
}

func (ll *Stack) isEmpty() bool {
	return ll.length == 0
}

func (ll *Stack) len() int {
	return ll.length
}
