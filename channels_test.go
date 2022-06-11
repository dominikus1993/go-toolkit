package gotoolkit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func rangeInt(from, to int) []int {
	res := make([]int, 0, to-from)
	for i := from; i < to; i++ {
		res = append(res, i)
	}
	return res
}

// BenchmarkToSlice-4       1302279               897.8 ns/op           688 B/op         12 allocs/op
func TestToSlice(t *testing.T) {
	numbers := make(chan int, 10)
	go func() {
		for _, a := range rangeInt(0, 10) {
			numbers <- a
		}
		close(numbers)
	}()

	result := ToSlice(numbers)
	assert.Len(t, result, 10)
	assert.ElementsMatch(t, rangeInt(0, 10), result)
}

func BenchmarkToSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numbers := make(chan int, 10)
		go func() {
			for _, a := range rangeInt(0, 10) {
				numbers <- a
			}
			close(numbers)
		}()

		ToSlice(numbers)
	}
}

func TestFanIn(t *testing.T) {
	numbers := FromSlice(rangeInt(1, 10))
	numbers2 := FromSlice(rangeInt(10, 20))

	result := FanIn(context.TODO(), numbers, numbers2)
	subject := ToSlice(result)
	assert.Len(t, subject, 19)
	assert.ElementsMatch(t, rangeInt(1, 20), subject)
}

func BenchmarkFanIn(b *testing.B) {
	ctx := context.TODO()
	for n := 0; n < b.N; n++ {
		numbers := FromSlice(rangeInt(1, 10))
		numbers2 := FromSlice(rangeInt(10, 20))

		result := FanIn(ctx, numbers, numbers2)
		ToSlice(result)
	}
}
