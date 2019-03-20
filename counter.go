package num

import (
	"sync"
)

type Counter struct {
	value uint
	max   uint
	mx    *sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
		max:   MaxUint(),
		mx:    new(sync.RWMutex),
	}
}

func (c *Counter) Value() uint {
	c.mx.RLock()
	defer c.mx.RUnlock()

	return c.value
}

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

func (c *Counter) Increment() {
	c.mx.Lock()
	defer c.mx.Unlock()

	if c.value+1 < c.max {
		c.value++
	} else {
		c.value = 0
	}
}
