package taskd7

func Mirror[T any](arr []T) []T {
	m := make([]T, len(arr))
	for i, v := range arr {
		m[len(arr)-i-1] = v
	}
	return m
}