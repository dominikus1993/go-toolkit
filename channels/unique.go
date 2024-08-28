package channels

func Uniq[T comparable](s <-chan T, size int) <-chan T {
	res := make(chan T, size)
	go func() {
		defer close(res)
		seen := make(map[T]bool)
		for v := range s {
			if _, ok := seen[v]; !ok {
				seen[v] = true
				res <- v
			}
		}
	}()
	return res
}

func UniqBy[T any, TKey comparable](s <-chan T, f func(T) TKey, size int) <-chan T {
	res := make(chan T, size)
	go func() {
		defer close(res)
		seen := make(map[TKey]bool)
		for v := range s {
			key := f(v)
			if _, ok := seen[key]; !ok {
				seen[key] = true
				res <- v
			}
		}
	}()
	return res
}
