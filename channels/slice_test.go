package channels

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkToSlice-4       1302279               897.8 ns/op           688 B/op         12 allocs/op
func TestToSlice(t *testing.T) {
	numbers := make(chan int, 10)
	go func() {
		for _, a := range RangeInt(0, 10) {
			numbers <- a
		}
		close(numbers)
	}()

	result := ToSlice(numbers)
	assert.Len(t, result, 10)
	assert.ElementsMatch(t, RangeInt(0, 10), result)
}

func TestFromSeq(t *testing.T) {
	seqOfInt := IntSeq(0, 10)
	numbers := FromSeq(seqOfInt)

	result := ToSlice(numbers)
	assert.Len(t, result, 10)
	assert.ElementsMatch(t, RangeInt(0, 10), result)
}

func TestMapFromSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4}

	result := MapFromSlice(numbers, func(el int) int { return el * 2 }, 10)

	subject := ToSlice(result)
	assert.Len(t, subject, len(numbers))
	assert.ElementsMatch(t, []int{2, 4, 6, 8}, subject)
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	numbers := FromSlice(arr)

	result := Map(numbers, func(el int) int { return el * 2 }, 10)

	subject := ToSlice(result)
	assert.Len(t, subject, len(arr))
	assert.ElementsMatch(t, []int{2, 4, 6, 8}, subject)
}
func BenchmarkToSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numbers := make(chan int, 10)
		go func() {
			for _, a := range RangeInt(0, 10) {
				numbers <- a
			}
			close(numbers)
		}()

		_ = ToSlice(numbers)
	}
}

func TestFanIn(t *testing.T) {
	numbers := FromSlice(RangeInt(1, 10))
	numbers2 := FromSlice(RangeInt(10, 20))

	result := FanIn(context.TODO(), numbers, numbers2)
	subject := ToSlice(result)
	assert.Len(t, subject, 19)
	assert.ElementsMatch(t, RangeInt(1, 20), subject)
}

func BenchmarkFanIn(b *testing.B) {
	ctx := context.TODO()
	for n := 0; n < b.N; n++ {
		numbers := FromSlice(RangeInt(1, 10))
		numbers2 := FromSlice(RangeInt(10, 20))

		result := FanIn(ctx, numbers, numbers2)
		ToSlice(result)
	}
}

func TestFilter(t *testing.T) {
	numbers := FromSlice(RangeInt(1, 10))
	result := Filter(numbers, func(element int) bool { return element%2 == 0 }, 10)
	subject := ToSlice(result)
	assert.Len(t, subject, 4)
	assert.ElementsMatch(t, []int{2, 4, 6, 8}, subject)
}

func BenchmarkFilter(b *testing.B) {
	arr := RangeInt(1, 10)
	for n := 0; n < b.N; n++ {
		numbers := FromSlice(arr)
		_ = ToSlice(Filter(numbers, func(element int) bool { return element%2 == 0 }, 1))
	}
}
