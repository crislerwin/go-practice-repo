package hashmap_test

import (
	"testing"

	"github.com/crislerwin/godsa/data-structures/hashmap"
)

func TestHashmap(t *testing.T) {
	mp := hashmap.New()
	t.Run("Test 1: Put(10) and checking if Get() is correct", func(t *testing.T) {
		mp.Put("test", 10)
		got := mp.Get("test")
		if got != 10 {
			t.Errorf("Put: %v Got: %v", 10, got)
		}
	})
	t.Run("Test 2: Reassigning the value and checking if Get() is correct", func(t *testing.T) {
		mp.Put("test", 20)
		got := mp.Get("test")
		if got != 20 {
			t.Errorf("Put (reassign): %v Got: %v", 20, got)
		}
	})
	t.Run("Test 3: Adding  new key when there is already some data", func(t *testing.T) {
		mp.Put("test2", 30)
		got := mp.Get("test2")
		if got != 30 {
			t.Errorf("Put (new key) %v Got: %v", 30, got)
		}
	})
}
