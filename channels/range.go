package channels

import "iter"

func RangeInt(fromInclusive, toExclusive int) []int {
	res := make([]int, 0, toExclusive-fromInclusive)
	for i := fromInclusive; i < toExclusive; i++ {
		res = append(res, i)
	}
	return res
}

func IntSeq(fromInclusive, toExclusive int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := fromInclusive; i < toExclusive; i++ {
			if !yield(i) {
				return
			}
		}
	}
}
