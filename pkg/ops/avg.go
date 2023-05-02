package ops

func Avg[T ~int | ~float64]() func(x T) T {
	sum, count := T(0), T(0)
	return func(x T) T {
		sum += x
		count++
		return sum / count
	}
}
