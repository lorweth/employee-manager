package ops

func Avg[T ~int | ~uint | float64 | float32]() func(x T) T {
	sum, count := T(0), T(0)
	return func(x T) T {
		sum += x
		count++
		return sum / count
	}
}
