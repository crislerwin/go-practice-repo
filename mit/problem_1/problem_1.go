package problem1

import "github.com/crislerwin/godsa/math/constraints"

func countLongSubArray(A []int) int {
	count := 0
	n := len(A)

	for i := 0; i < n; i++ {
		minVal := A[i]
		maxVal := A[i]

		for j := i; j < n; j++ {
			minVal = Min(minVal, A[j])
			maxVal = Max(maxVal, A[j])

			if maxVal-minVal >= j-i {
				count++
			}
		}
	}

	return count
}

func Min[T constraints.Integer](values ...T) T {
	minVal := values[0]
	for _, value := range values {
		if value < minVal {
			minVal = value
		}
	}
	return minVal
}

func Max[T constraints.Integer](values ...T) T {
	maxVal := values[0]
	for _, value := range values {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}
