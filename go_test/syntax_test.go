package go_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
	v1, ok := <- ch1
	assert.Equal(t, 1, v1)
	assert.Equal(t, true, ok)

	ch2 := make(chan int, 1)
	ch2 <- 1
	close(ch2)
	v2, ok := <- ch2
	assert.Equal(t, 1, v2)
	assert.Equal(t, true, ok)
}



func TestVariadic(t *testing.T) {
	f := func (a int, b ...int) int{
		return len(b)
	}
	assert.Equal(t, 0, f(1))
	assert.Equal(t, 2, f(1,2,3))
	var b []int
	assert.Nil(t, b)
	b = make([]int, 0)
	assert.Equal(t, 0, f(1, b... ))
	b = append(b, 1, 2, 3)
	assert.Equal(t, 3, f(1, b... ))
}