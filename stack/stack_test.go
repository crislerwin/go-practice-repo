package stack

import (
	"reflect"
	"testing"
)

func TestStackArray(t *testing.T) {
	t.Run("Stack With Array", func(t *testing.T) {
		stackPush(2)
		stackPush(3)
		t.Run("Stack Push", func(t *testing.T) {
			if !reflect.DeepEqual([]any{3, 2}, stackArray) {
				t.Errorf("Stack Push is not work we expected %v but got %v", []any{3, 2}, stackArray)
			}
		})

		pop := stackPop()

		t.Run("Stack Pop", func(t *testing.T) {
			if stackLength() == 2 && pop != 3 {
				t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
			}
		})
		stackPush(2)
		stackPush(83)
		t.Run("Stack Peak", func(t *testing.T) {
			if stackPeak() != 83 {
				t.Errorf("Stack Peak is not work we expected %v but got %v", 3, stackLength())
			}
		})

		t.Run("Stack Length", func(t *testing.T) {
			if stackLength() != 3 {
				t.Errorf("Stack Length is not work we expected %v but got %v", 3, stackLength())
			}
		})

		t.Run("Stack Empty", func(t *testing.T) {
			if stackEmpty() == true {
				t.Errorf("Stack Empty is not working we expected %v but got %v", true, stackEmpty())
			}
		})
		stackPop()
		stackPop()
		stackPop()

		if stackEmpty() == false {
			t.Errorf("Stack empty is not work we expected %v but got %v", true, stackEmpty())
		}
	})
}

func TestStackLinkedList(t *testing.T) {
	var newStack Stack

	newStack.push(1)
	newStack.push(2)
	t.Run("Stack Push", func(t *testing.T) {
		result := newStack.show()
		expected := []any{2, 1}
		for x := range result {
			if result[x] != expected[x] {
				t.Errorf("Stack Push is not work, got %v but expected %v", result, expected)
			}
		}
	})
	t.Run("Stack isEmpty", func(t *testing.T) {
		if newStack.isEmpty() {
			t.Error("Stack Length is returned true but expected false", newStack.isEmpty())
		}
	})
	t.Run("Stack Length", func(t *testing.T) {
		if newStack.len() != 2 {
			t.Error("Stack length should be 2 but got", newStack.len())
		}
	})
}
