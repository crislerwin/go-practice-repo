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

	t.Run("Test AddAtBeg()", func(t *testing.T) {
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

	list.AddAtEnd(4)
	t.Run("Test AddAtEnd()", func(t *testing.T) {
		want := []any{3, 2, 1, 4}
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

	t.Run("Test DelAtBeg()", func(t *testing.T) {
		want := any(3)
		got, ok := list.DelAtBeg()
		if !ok {
			t.Error("unexpected not ok")
		}
		if got != want {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtEnd()", func(t *testing.T) {
		want := any(4)
		got, ok := list.DelAtEnd()
		if !ok {
			t.Error("unexpected not ok")
		}
		if got != want {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})

}
