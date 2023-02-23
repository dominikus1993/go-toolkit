package channels

func RangeInt(fromInclusive, toExclusive int) []int {
	res := make([]int, 0, toExclusive-fromInclusive)
	for i := fromInclusive; i < toExclusive; i++ {
		res = append(res, i)
	}
	return res
}
