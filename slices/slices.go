package slices

func Filter[S ~[]T, T any](s S, f func(T) bool) []T {
	filterResult := make(S, 0, len(s))
	for _, v := range s {
		if f(v) {
			filterResult = append(filterResult, v)
		}
	}
	return filterResult
}
