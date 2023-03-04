package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	ID   string
	Name string
}

func TestUnique(t *testing.T) {
	numbers := make(chan int, 10)
	go func() {
		for _, a := range []int{1, 1, 2, 2, 3, 3, 4, 4} {
			numbers <- a
		}
		close(numbers)
	}()

	result := ToSlice(Uniq(numbers, 10))
	assert.Len(t, result, 4)
	assert.ElementsMatch(t, []int{1, 2, 3, 4}, result)
}

func TestUniqueBy(t *testing.T) {
	numbers := make(chan testData, 10)
	go func() {
		elements := []testData{
			{ID: "2137", Name: "tets"},
			{ID: "2137", Name: "tets"},
			{ID: "213769", Name: "tets"},
			{ID: "213769", Name: "tets"},
		}
		for _, a := range elements {
			numbers <- a
		}
		close(numbers)
	}()

	result := ToSlice(UniqBy(numbers, func(elem testData) string { return elem.ID }, 10))
	assert.Len(t, result, 2)
	assert.ElementsMatch(t, []testData{
		{ID: "2137", Name: "tets"},
		{ID: "213769", Name: "tets"},
	}, result)
}

func BenchmarkUnique(b *testing.B) {
	numbers := make(chan int, 10)
	go func() {
		for _, a := range []int{1, 1, 2, 2, 3, 3, 4, 4} {
			numbers <- a
		}
		close(numbers)
	}()

	for i := 0; i < b.N; i++ {
		_ = ToSlice(Uniq(numbers, 10))
	}

}

func BenchmarkUniqueBy(b *testing.B) {
	numbers := make(chan testData, 10)
	go func() {
		elements := []testData{
			{ID: "2137", Name: "tets"},
			{ID: "2137", Name: "tets"},
			{ID: "213769", Name: "tets"},
			{ID: "213769", Name: "tets"},
		}
		for _, a := range elements {
			numbers <- a
		}
		close(numbers)
	}()

	for i := 0; i < b.N; i++ {
		_ = ToSlice(UniqBy(numbers, func(elem testData) string { return elem.ID }, 10))
	}
}
