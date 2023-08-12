package linkedlist

import (
	"reflect"
	"testing"
)

func TestSingly(t *testing.T) {
	list := NewSingly[int]()
	list.AddAtBeg(1)
	list.AddAtBeg(2)
	list.AddAtBeg(3)

	t.Run("Test Add At Beg", func(t *testing.T) {
		want := []any{3, 2, 1}
		got := []any{}
		current := list.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)

		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})
}
