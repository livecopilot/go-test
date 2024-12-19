package study_map_reduce

func Map[T any, R any](input []T, f func(T) R) []R {
	result := make([]R, len(input)) // 预分配长度
	for i := range input {
		result[i] = f(input[i]) // 避免额外变量拷贝
	}
	return result
}

func Filter[T any](input []T, predicate func(T) bool) []T {
	count := 0
	for _, v := range input {
		if predicate(v) {
			count++ // 预计算满足条件的元素数量
		}
	}
	result := make([]T, 0, count) // 预分配容量
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T any](input []T, f func(T, T) T, initial T) T {
	if len(input) == 0 {
		return initial
	}
	
	result := initial
	for _, v := range input {
		result = f(result, v)
	}
	return result
}
