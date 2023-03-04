package random

import (
	"testing"

	"github.com/dominikus1993/go-toolkit/channels"
	"github.com/stretchr/testify/assert"
)

func TestTakeRandom(t *testing.T) {
	numbers := make(chan int, 10)
	go func() {
		for _, a := range channels.RangeInt(0, 10) {
			numbers <- a
		}
		close(numbers)
	}()

	subject := channels.ToSlice(TakeRandom(numbers, 2))

	assert.Len(t, subject, 2)
}

func TestTakeRandomTwice(t *testing.T) {
	numbers := channels.RangeInt(0, 10)

	subject := channels.ToSlice(TakeRandom(channels.FromSlice(numbers), 2))
	subject2 := channels.ToSlice(TakeRandom(channels.FromSlice(numbers), 2))
	assert.Len(t, subject, 2)
	assert.Len(t, subject2, 2)

	assert.NotEqual(t, subject, subject2)
}
