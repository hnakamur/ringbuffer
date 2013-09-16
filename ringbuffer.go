package ringbuffer

import (
	"errors"
)

type RingBuffer struct {
	buffer   []interface{}
	start    int
	end      int
}

func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		buffer:   make([]interface{}, capacity),
		start:    0,
		end:      0,
	}
}

func (b *RingBuffer) Capacity() int {
	return len(b.buffer)
}

func (b *RingBuffer) Len() int {
	return b.end - b.start
}

func (b *RingBuffer) Add(item interface{}) error {
	if b.Len() >= len(b.buffer) {
		return errors.New("buffer full")
	}
	b.buffer[b.end % len(b.buffer)] = item
	b.end++
	return nil
}

func (b *RingBuffer) Remove() (interface{}, error) {
	if b.Len() <= 0 {
		return nil, errors.New("buffer empty")
	}
	item := b.buffer[b.start % len(b.buffer)]
	b.start++
	return item, nil
}
