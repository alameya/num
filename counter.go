package num

import (
	"sync"
)

type Counter struct {
	value int
	max   int
	mx    *sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
		max:   MaxInt(),
		mx:    new(sync.RWMutex),
	}
}

func (c *Counter) Value() int {
	c.mx.RLock()
	defer c.mx.RUnlock()

	return c.value
}

func (c *Counter) Max(maxValue int) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if maxValue > MaxInt() {
		maxValue = MaxInt()
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
