package random

import (
	"math/rand"
	"time"

	"github.com/dominikus1993/go-toolkit/channels"
)

var randWithSeed *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func TakeRandomFromChannel[T any](s <-chan T, take int) <-chan T {
	res := make(chan T, take)
	go func() {
		defer close(res)
		slice := channels.ToSlice(s)
		sliceLength := len(slice)
		var takeNum = take
		if take >= sliceLength {
			takeNum = sliceLength
		}
		for i := 0; i < takeNum; i++ {
			index := randWithSeed.Intn(sliceLength)
			res <- slice[index]
		}
	}()
	return res
}

func TakeRandomFromSlice[T any](slice []T, take int) []T {
	if take == 0 {
		return make([]T, 0)
	}
	sliceLength := len(slice)
	if take >= sliceLength {
		return slice
	}
	randomIndexes := generateUniqueRandomNumbers(take, sliceLength)
	randomArticles := make([]T, 0, take)
	for index := range randomIndexes {
		randomArticles = append(randomArticles, slice[index])
	}

	return randomArticles
}

func TakeRandomToSlice[T any](s <-chan T, take int) []T {
	if take == 0 {
		return make([]T, 0)
	}
	slice := channels.ToSlice(s)
	sliceLength := len(slice)
	if take >= sliceLength {
		return slice
	}
	randomIndexes := generateUniqueRandomNumbers(take, sliceLength)
	randomArticles := make([]T, 0, take)
	for index := range randomIndexes {
		randomArticles = append(randomArticles, slice[index])
	}

	return randomArticles
}

func generateUniqueRandomNumbers(n, max int) []int {
	set := make(map[int]bool)
	var result []int
	for len(set) < n {
		value := randWithSeed.Intn(max)
		if !set[value] {
			set[value] = true
			result = append(result, value)
		}
	}
	return result
}
