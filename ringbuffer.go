package ringbuffer

import (
	"errors"
)

type RingBuffer struct {
	buffer   []interface{}
	capacity int
	mask     int
	start    int
	end      int
}

func NewRingBuffer(capacity int) *RingBuffer {
	bufSize := smallestPow2(capacity)
	return &RingBuffer{
		buffer:   make([]interface{}, bufSize),
		capacity: capacity,
		mask:     bufSize - 1,
		start:    0,
		end:      0,
	}
}

func (b *RingBuffer) Capacity() int {
	return b.capacity
}

func (b *RingBuffer) Len() int {
	return b.end - b.start
}

func (b *RingBuffer) Add(item interface{}) error {
	if b.Len() >= b.capacity {
		return errors.New("buffer full")
	}
	b.buffer[b.end&b.mask] = item
	b.end++
	return nil
}

func (b *RingBuffer) Remove() (interface{}, error) {
	if b.Len() <= 0 {
		return nil, errors.New("buffer empty")
	}
	item := b.buffer[b.start&b.mask]
	b.start++
	return item, nil
}

func smallestPow2(n int) int {
	var x int
	for x = 1; x < n; x <<= 1 {
	}
	return x
}
