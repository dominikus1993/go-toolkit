package random

import (
	"math/rand"
	"time"
)

var randWithSeed *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func TakeRandom[T any](s <-chan T, take int) <-chan T {
	res := make(chan T, take)
	go func() {
		defer close(res)
		itemsStreamed := 0
		for v := range s {
			if itemsStreamed >= take {
				return
			}
			num := randWithSeed.Int()
			if num%2 == 0 {
				itemsStreamed = itemsStreamed + 1
				res <- v
			}
		}
		close(res)
	}()
	return res
}
