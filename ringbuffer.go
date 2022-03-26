package ringbuffer

import "io"

// RingBuffer is a generic circular buffer that can hold any type of data.
// The PushBack method does not fail in case the buffer is full, it instead overwrites
// the oldest data slot and moves the read index forward.
type RingBuffer[A any] struct {
	b                   []A
	nextWrite, nextRead int
	len                 int
}

// NewRingBuffer returns a new RingBuffer with specified size and type.
func NewRingBuffer[A any](size int) *RingBuffer[A] {
	return &RingBuffer[A]{
		b: make([]A, size),
	}
}

// Cap returns the capacity of the RingBuffer. That is the specified size.
func (r *RingBuffer[A]) Cap() int {
	return len(r.b)
}

// Len returns the number of valid data slots available to read
// in the RingBuffer.
func (r *RingBuffer[A]) Len() int {
	return r.len
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// PushBack writes a new element in the RingBuffer. If the buffer is full,
// it overwrites the oldest element and moves forward the read index.
func (r *RingBuffer[A]) PushBack(elem A) {
	r.b[r.nextWrite] = elem
	if r.nextWrite == r.nextRead {
		r.nextRead = (r.nextRead + 1) % len(r.b)
	}
	r.nextWrite = (r.nextWrite + 1) % len(r.b)
	r.len = min(r.Cap(), r.len+1)
	return
}

// PopFront reads the oldest element from the buffer.
// It returns io.EOF error in case the buffer is empty.
func (r *RingBuffer[A]) PopFront() (A, error) {
	if r.len == 0 {
		return *new(A), io.EOF
	}
	ret := r.b[r.nextRead]
	r.nextRead = (r.nextRead + 1) % len(r.b)
	r.len--
	return ret, nil
}
