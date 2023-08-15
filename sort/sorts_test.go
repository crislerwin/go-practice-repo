package sort_test

import (
	"reflect"
	"testing"

	"github.com/crislerwin/go-practice/sort"
)

func testFramework(t *testing.T, sortFn func([]int) []int) {
	sortTests := []struct {
		input    []int
		expected []int
		name     string
	}{{
		input:    []int{1, 2, 3, 5, 4, 6, 7, 10, 9, 8},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:     "Sorted Unsigned",
	}}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := sortFn(test.input)
			sorted := reflect.DeepEqual(actual, test.expected)
			if !sorted {
				t.Errorf("test %s failed", test.name)
				t.Errorf("Actual %v, expected: %v", actual, test.expected)
			}
		})
	}
}
func TestSort(t *testing.T) {
	testFramework(t, sort.Bubble[int])
}
