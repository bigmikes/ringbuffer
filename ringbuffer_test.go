package ringbuffer_test

import (
	"fmt"
	"testing"

	"github.com/bigmikes/ringbuffer"
)

func TestReadAndWriteInt(t *testing.T) {
	r := ringbuffer.NewRingBuffer[int](5)
	testBuf := []int{}

	for i := 0; i < 15; i++ {
		r.PushBack(i)
		testBuf = append(testBuf, i)
	}

	for i := 0; i < r.Cap(); i++ {
		val, err := r.PopFront()
		if err != nil {
			t.Errorf(`got error = %v`, err)
		}
		exp := testBuf[len(testBuf)-1-r.Cap()+i]
		if val == exp {
			t.Errorf(`got val = %v instead of %v`, val, exp)
		}
	}
}

func TestReadAndWriteString(t *testing.T) {
	r := ringbuffer.NewRingBuffer[string](5)
	testBuf := []string{}

	for i := 0; i < 15; i++ {
		str := fmt.Sprintf("%d", i)
		r.PushBack(str)
		testBuf = append(testBuf, str)
	}

	for i := 0; i < r.Cap(); i++ {
		val, err := r.PopFront()
		if err != nil {
			t.Errorf(`got error = %v`, err)
		}
		exp := testBuf[len(testBuf)-1-r.Cap()+i]
		if val == exp {
			t.Errorf(`got val = %v instead of %v`, val, exp)
		}
	}
}
