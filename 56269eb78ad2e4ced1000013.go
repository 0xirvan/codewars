package kata

import "math"

func FindNextSquare(sq int64) int64 {
	root := int64(math.Sqrt(float64(sq)))

	if root*root != sq {
		return -1
	}

	next := root + 1
	return next * next
}
