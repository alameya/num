package num

import (
	"math"
	"math/bits"
)

func MaxInt() int {
	if bits.UintSize == 32 {
		return math.MaxInt32
	}

	return  math.MaxInt64
}

func MinInt() int {
	if bits.UintSize == 32 {
		return math.MinInt32
	}

	return math.MinInt64
}