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
}
