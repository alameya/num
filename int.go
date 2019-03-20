package num

import (
	"math"
	"math/bits"
)

func MaxUint() uint {
	if bits.UintSize == 64 {
		return  math.MaxUint64
	}

	return math.MaxUint32
}
