package ringbuffer

import (
	"testing"
)

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
