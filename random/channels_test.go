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

	subject := channels.ToSlice(TakeRandomFromChannel(numbers, 2))

	assert.Len(t, subject, 2)
}

func TestTakeRandomWhenStreamHasSmallerElementsQuantityThanTakeParam(t *testing.T) {
	numbers := make(chan int, 10)
	go func() {
		for _, a := range channels.RangeInt(0, 10) {
			numbers <- a
		}
		close(numbers)
	}()

	subject := channels.ToSlice(TakeRandomFromChannel(numbers, 20))

	assert.Len(t, subject, 10)
}

func TestTakeRandomTwice(t *testing.T) {
	numbers := channels.RangeInt(0, 10)

	subject := channels.ToSlice(TakeRandomFromChannel(channels.FromSlice(numbers), 2))
	subject2 := channels.ToSlice(TakeRandomFromChannel(channels.FromSlice(numbers), 2))
	assert.Len(t, subject, 2)
	assert.Len(t, subject2, 2)

	assert.NotEqual(t, subject, subject2)
}

func TestTakeRandomToSlice(t *testing.T) {
	numbers := make(chan int, 10)
	go func() {
		for _, a := range channels.RangeInt(0, 10) {
			numbers <- a
		}
		close(numbers)
	}()

	subject := TakeRandomToSlice(numbers, 2)

	assert.Len(t, subject, 2)
}

func TestTakeRandomWhenStreamHasSmallerElementsQuantityThanTakeParamToSlice(t *testing.T) {
	numbers := make(chan int, 10)
	go func() {
		for _, a := range channels.RangeInt(0, 10) {
			numbers <- a
		}
		close(numbers)
	}()

	subject := TakeRandomToSlice(numbers, 20)

	assert.Len(t, subject, 10)
}

func TestTakeRandomTwiceToSlice(t *testing.T) {
	numbers := channels.RangeInt(0, 10)

	subject := TakeRandomToSlice(channels.FromSlice(numbers), 2)
	subject2 := TakeRandomToSlice(channels.FromSlice(numbers), 2)
	assert.Len(t, subject, 2)
	assert.Len(t, subject2, 2)

	assert.NotEqual(t, subject, subject2)
}
