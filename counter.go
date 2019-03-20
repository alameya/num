package num

import (
	"sync"
)

// Counter thread safe Increment() calls calculator.
// Use NewCounter for creation.
// If current value equal maximum set counter to 0.
type Counter struct {
	value uint
	max   uint
	mx    *sync.RWMutex
}

// NewCounter return a new Counter with default values.
func NewCounter() *Counter {
	return &Counter{
		max:   MaxUint(),
		mx:    new(sync.RWMutex),
	}
}

// Value return current counter value.
func (c *Counter) Value() uint {
	c.mx.RLock()
	defer c.mx.RUnlock()

	return c.value
}

// Max set max value.
// Drop down counter to 0 if less then current value.
func (c *Counter) Max(maxValue uint) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if maxValue > MaxUint() {
		maxValue = MaxUint()
	}

	c.max = maxValue

	if c.value >= c.max {
		c.value = 0
	}
}

// Increment increase current value by 1.
// If next value is maximum drop down current value to 0.
func (c *Counter) Increment() {
	c.mx.Lock()
	defer c.mx.Unlock()

	if c.value+1 < c.max {
		c.value++
	} else {
		c.value = 0
	}
}
