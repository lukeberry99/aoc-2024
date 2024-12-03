package slices

import "sort"

func SlidingWindow(slice []int, size int) [][]int {
	var result [][]int = make([][]int, 0)
	for i := 0; i <= len(slice)-size; i++ {
		window := slice[i : i+size]
		result = append(result, window)
	}
	return result
}

func SortSlice(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return slice
}
