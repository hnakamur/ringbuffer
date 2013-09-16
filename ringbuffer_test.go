package ringbuffer

import (
	"testing"
)

func TestNewRingBuffer(t *testing.T) {
	buf := NewRingBuffer(3)
	if buf.Capacity() != 3 {
		t.Errorf("capacity: exptected 3, got %d", buf.Capacity())
	}
}

func TestAdd(t *testing.T) {
	buf := NewRingBuffer(3)

	err := buf.Add(1)
	if err != nil {
		t.Error(err)
	}
	if buf.Len() != 1 {
		t.Errorf("len: exptected 1, got %d", buf.Len())
	}

	err = buf.Add(2)
	if err != nil {
		t.Error(err)
	}
	if buf.Len() != 2 {
		t.Errorf("len: exptected 2, got %d", buf.Len())
	}

	err = buf.Add(3)
	if err != nil {
		t.Error(err)
	}
	if buf.Len() != 3 {
		t.Errorf("len: exptected 3, got %d", buf.Len())
	}

	err = buf.Add(4)
	if err == nil {
		t.Error("add: expected an error")
	}
	if err.Error() != "buffer full" {
		t.Error("add: expected buffer full error")
	}
}

func TestRemove(t *testing.T) {
	buf := NewRingBuffer(3)

	err := buf.Add(1)
	if err != nil {
		t.Error(err)
	}
	err = buf.Add(2)
	if err != nil {
		t.Error(err)
	}
	err = buf.Add(3)
	if err != nil {
		t.Error(err)
	}

	item, err := buf.Remove()
	if err != nil {
		t.Error(err)
	}
	if item != 1 {
		t.Errorf("remove: exptected 1, got %d", item)
	}

	item, err = buf.Remove()
	if err != nil {
		t.Error(err)
	}
	if item != 2 {
		t.Errorf("remove: exptected 2, got %d", item)
	}

	err = buf.Add(4)
	if err != nil {
		t.Error(err)
	}

	item, err = buf.Remove()
	if err != nil {
		t.Error(err)
	}
	if item != 3 {
		t.Errorf("remove: exptected 3, got %d", item)
	}

	item, err = buf.Remove()
	if err != nil {
		t.Error(err)
	}
	if item != 4 {
		t.Errorf("remove: exptected 4, got %d", item)
	}

	item, err = buf.Remove()
	if err == nil {
		t.Error("remove: expected an error")
	}
	if err.Error() != "buffer empty" {
		t.Error("remove: expected buffer empty error")
	}

	err = buf.Add(5)
	if err != nil {
		t.Error(err)
	}

	err = buf.Add(6)
	if err != nil {
		t.Error(err)
	}

	item, err = buf.Remove()
	if err != nil {
		t.Error(err)
	}
	if item != 5 {
		t.Errorf("remove: exptected 5, got %d", item)
	}

	item, err = buf.Remove()
	if err != nil {
		t.Error(err)
	}
	if item != 6 {
		t.Errorf("remove: exptected 6, got %d", item)
	}

	if buf.Len() != 0 {
		t.Errorf("len: exptected 0, got %d", buf.Len())
	}
}

func TestSmalltestPow2(t *testing.T) {
	tests := []struct{
		n int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 4},
		{5, 8},
		{6, 8},
		{7, 8},
		{8, 8},
		{12, 16},
		{1025, 2048},
	}
	for _, test := range tests {
		actual := smallestPow2(test.n)
		if actual != test.expected {
			t.Errorf("got %d, expected %d for %d", actual, test.expected,
				test.n)
		}
	}
}
