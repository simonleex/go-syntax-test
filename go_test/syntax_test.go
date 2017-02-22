package go_test

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestClosedChannel(t *testing.T) {
	ch := make(chan int, 1)
	//ch <- 1
	close(ch)
	v, ok := <-ch
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, v)

	ch1 := make(chan int, 1)
	ch1 <- 1
	v1, ok := <-ch1
	assert.Equal(t, 1, v1)
	assert.Equal(t, true, ok)

	ch2 := make(chan int, 1)
	ch2 <- 1
	close(ch2)
	v2, ok := <-ch2
	assert.Equal(t, 1, v2)
	assert.Equal(t, true, ok)
}

func TestVariadic(t *testing.T) {
	f := func(a int, b ...int) int {
		return len(b)
	}
	assert.Equal(t, 0, f(1))
	assert.Equal(t, 2, f(1, 2, 3))
	var b []int
	assert.Nil(t, b)
	b = make([]int, 0)
	assert.Equal(t, 0, f(1, b...))
	b = append(b, 1, 2, 3)
	assert.Equal(t, 3, f(1, b...))
}

func TestDeferInInnerFun(t *testing.T) {
	var lock sync.Mutex
	te := func() {
		lock.Lock()
		defer lock.Unlock()
	}
	te()
	lock.Lock()
	lock.Unlock()
}

func TestSliceAppendBasic(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3}
	s3 := append(s1[:1], s2...)
	assert.Equal(t, s3, s1)
	assert.Equal(t, []int{1, 1, 2, 3}, s1)
	s4 := append(s1[:2], s2...)
	assert.NotEqual(t, s4, s1)
	assert.Equal(t, s1, s3)
	assert.Equal(t, []int{1, 1, 2, 3}, s1)
	assert.Equal(t, []int{1, 1, 1, 2, 3}, s4)
}

func TestSliceAppend(t *testing.T) {
	f1 := func() []int {
		s1 := []int{1, 2, 3, 4}
		s2 := []int{-1, -2, -3}
		return append(append(s1[:1], s2...), s1[1:]...)
	}

	f2 := func() []int {
		s1 := []int{1, 2, 3, 4}
		s2 := []int{-1, -2, -3, -4}
		return append(append(s1[:1], s2...), s1[1:]...)
	}

	assert.Equal(t, []int{1, -1, -2, -3, -1, -2, -3}, f1())

	assert.Equal(t, []int{1, -1, -2, -3, -4, 2, 3, 4}, f2())
}

func TestNew(t *testing.T) {
	type A struct {
		a interface{}
	}
	C := new(A)
	assert.Equal(t, &A{nil}, C)

}

func TestDefer(t *testing.T) {
	c := func() (i int) {
		i = 0
		defer func() {
			i++
		}()
		return
	}

	assert.Equal(t, 1, c())

	c1 := func() int {
		i := 0
		defer func() {
			i++
		}()
		return i
	}

	assert.Equal(t, 0, c1())

	c2 := func() int {
		i := 0
		defer func(i int) {
			i++
		}(i)
		return i
	}

	assert.Equal(t, 0, c2())

	c3 := func() (i int) {
		i = 0
		defer func(i int) {
			i++
		}(i)
		return
	}

	assert.Equal(t, 0, c3())
}

func TestSliceCap(t *testing.T)  {
	sl := make([]int, 5, 6)
	assert.Equal(t, 5, len(sl))
	assert.Equal(t, 6, cap(sl))
	sl1 := []int{1}
	assert.Equal(t, 1, len(sl1))
	assert.Equal(t, 1, len(sl1))
}
