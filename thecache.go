// https://www.linkedin.com/learning/advanced-go-programming-data-structures-code-architecture-and-testing/code-challenges/urn:li:la_assessmentV2:50886190?autoSkip=true&resume=false
// Write your answer here, and then test your code.
// Your job is to implement the Get and Set Cache methods.

package main

import (
	"errors"
	"sync"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type Key string
type Value string

type ValueWriter interface {
	Set(k Key, v Value)
}
type ValueReader interface {
	Get(k Key) (Value, error)
}

// Make any required changes to the Cache struct.
type Cache struct {
	lookup map[Key]Value
	lock   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{lookup: make(map[Key]Value)}
}

func (c *Cache) Set(k Key, v Value) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.lookup[k] = v
}

func (c *Cache) Get(k Key) (Value, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	v, ok := c.lookup[k]
	if !ok {
		return Value(""), errors.New("key not found")
	}
	return v, nil
}

func WriteValues(w ValueWriter, keys []Key, values []Value) {
	for i, k := range keys {
		w.Set(k, values[i])
	}
}

func ReadValues(r ValueReader, keys []Key) ([]Value, error) {
	values := make([]Value, len(keys))
	for i, k := range keys {
		v, err := r.Get(k)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}

	return values, nil
}
