package popularcache

import (
	"sync"
)

func New[T any]() API[T] {
	return &cache[T]{
		items: make(map[string]T),
		order: make(map[string]int),
	}
}

type cache[T any] struct {
	// items are mapped to their corresponding expiry
	mu sync.Mutex

	// items has the ID as the key and the Item as the value
	items map[string]T
	// order contains the ID as the key and the order as the value. order
	// uses a zero index based system i.e. First item in slice is order 0 not 1.
	order map[string]int
}

func (c *cache[T]) Add(id string, item T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[id] = item
	c.shiftOrderDown(1)
	// Add the item at the beginning
	c.order[id] = 0
}

func (c *cache[T]) Collect(id string) (T, bool) {
	// Ensure collect brings the item to the front of the slice
	// First shift every item down by one
	c.shiftOrderDown(1)
	// Then move first item to the first position
	c.order[id] = 0
	// Fill gaps to have incremental indexing
	c.fillGaps()

	item, ok := c.items[id]
	return item, ok
}

func (c *cache[T]) List() []T {
	c.mu.Lock()
	defer c.mu.Unlock()

	ls := make([]T, len(c.items))
	for key, order := range c.order {
		ls[order] = c.items[key]
	}

	return ls
}

// shiftOrderDown is unsafe for multithreading as it accesses memory without accessing the mutex. shiftOrderDown
// should only be used when mutex has been obtained.
func (c *cache[T]) shiftOrderDown(by int) {
	for id, index := range c.order {
		c.order[id] += index + by
	}
}

// fillGaps moves items up the order if there is a incremental gap
func (c *cache[T]) fillGaps() {
	var highestIndex int
	indexedOrder := make(map[int]string)
	for id, order := range c.order {
		indexedOrder[order] = id

		// Keep track of the lowest index number
		if order > highestIndex {
			highestIndex = order
		}
	}

	var currIndex int
	for i := 0; i <= highestIndex; i++ {
		id, ok := indexedOrder[i]
		if !ok {
			continue
		}

		c.order[id] = currIndex
		// Increment only when we add to the order map
		currIndex++
	}
}
