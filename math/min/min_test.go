package min

import "testing"

type TestCase struct {
	name    string
	base    int
	numbers []int
	min     int
}

func makeTestCases() []TestCase {
	var testCases = []TestCase{{
		"Minimum of [128, 127], min = 117", 8, []int{128, 127}, 127,
	}}
	return testCases
}

func TestMin(t *testing.T) {

	for _, test := range makeTestCases() {
		t.Run(test.name, func(t *testing.T) {
			actualMin := Min(test.numbers...)
			if actualMin != test.min {
				t.Errorf("Expected %d, got %d", test.min, actualMin)
			}
		})
	}

}
