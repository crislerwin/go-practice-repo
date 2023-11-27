package linkedlist

import (
	"reflect"
	"testing"
)

func TestDouble(t *testing.T) {
	newList := NewDoubly[int]()
	newList.AddAtBeg(1)
	newList.AddAtBeg(2)
	newList.AddAtBeg(3)
	t.Run("Test addAtBeg", func(t *testing.T) {
		wantNext := []int{3, 2, 1}
		wantPrev := []int{1, 2, 3}

		got := []int{}
		current := newList.Head.Next
		got = append(got, current.Val)
		for current.Next != newList.Head {
			current = current.Next
			got = append(got, current.Val)
		}

		if !reflect.DeepEqual(got, wantNext) {
			t.Errorf("Got: %v, want: %v", got, wantNext)
		}

		got = []int{}
		got = append(got, current.Val)
		for current.Prev != newList.Head {
			current = current.Prev
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, wantPrev) {
			t.Errorf("got: %v, want: %v", got, wantPrev)
		}
	})
	newList.AddAtEnd(4)
	t.Run("Test AddAtEnd", func(t *testing.T) {
		want := []int{3, 2, 1, 4}
		got := []int{}
		current := newList.Head.Next
		got = append(got, current.Val)
		for current.Next != newList.Head {
			current = current.Next
			got = append(got, current.Val)

		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})
	t.Run("Test DelAtBeg", func(t *testing.T) {
		want := 3
		got, ok := newList.DelAtBeg()
		if !ok {
			t.Error("unexpected not-ok")

		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

}
