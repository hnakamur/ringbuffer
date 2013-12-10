package ringbuffer

import (
	"errors"
)

type RingBuffer struct {
	buffer []interface{}
	start  int
	end    int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]interface{}, size+1),
		start:  0,
		end:    0,
	}
}

func (b *RingBuffer) Full() bool {
	return b.inc(b.end) == b.start
}

func (b *RingBuffer) Empty() bool {
	return b.end == b.start
}

func (b *RingBuffer) Len() int {
	l := b.end - b.start
	if l < 0 {
		l += len(b.buffer)
	}
	return l
}

func (b *RingBuffer) inc(index int) int {
	return (index + 1) % len(b.buffer)
}

func (b *RingBuffer) Add(item interface{}) error {
	if b.Full() {
		return errors.New("buffer full")
	}
	b.buffer[b.end] = item
	b.end = b.inc(b.end)
	return nil
}

func (b *RingBuffer) Remove() (interface{}, error) {
	if b.Empty() {
		return nil, errors.New("buffer empty")
	}
	item := b.buffer[b.start]
	b.start = b.inc(b.start)
	return item, nil
}
