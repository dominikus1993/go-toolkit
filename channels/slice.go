package channels

import (
	"context"
	"sync"
)

func FromSlice[T any](s []T) <-chan T {
	res := make(chan T)
	go func() {
		for _, v := range s {
			res <- v
		}
		close(res)
	}()
	return res
}

func ToSlice[T any](s <-chan T) []T {
	res := make([]T, 0)
	for v := range s {
		res = append(res, v)
	}
	return res
}

func MapFromSlice[T any, TMap any](s []T, fmap func(el T) TMap, size int) <-chan TMap {
	res := make(chan TMap, size)
	go func() {
		for _, v := range s {
			res <- fmap(v)
		}
		close(res)
	}()
	return res
}

func FanIn[T any](ctx context.Context, stream ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)
	output := func(c <-chan T, wait *sync.WaitGroup) {
		defer wait.Done()
		for v := range c {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}
	wg.Add(len(stream))
	for _, c := range stream {
		go output(c, &wg)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
