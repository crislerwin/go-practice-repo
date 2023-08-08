package problem1

import "testing"

func TestCountLongSubArray(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 4, 1, 4}, 13},
	}
	for i, test := range testCases {
		resut := countLongSubArray(test.input)
		if resut != test.expected {
			t.Errorf("Test %d: expected %d, got %d", i, test.expected, resut)
		}
	}

}
