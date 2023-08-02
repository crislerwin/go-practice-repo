package stack

import "container/list"

type SList struct {
	stack *list.List
}

func (sl *SList) Push(val any) {
	sl.stack.PushFront(val)
}

func (sl *SList) Length() int {
	return sl.stack.Len()
}
