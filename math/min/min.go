package min

import "github.com/crislerwin/godsa/math/constraints"

func Min[T constraints.Integer](values ...T) T {
	minVal := values[0]
	for _, value := range values {
		if value < minVal {
			minVal = value
		}
	}
	return minVal
}
