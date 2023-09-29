package driven

func Relu[T ~int | ~float64](val T) T {
	if val < 0 {
		return 0
	}
	return val
}
