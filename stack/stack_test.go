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
	})
}

func BenchmarkStackArrayPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stackPush(i)
	}
}
