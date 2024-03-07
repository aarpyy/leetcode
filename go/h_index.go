package leetcode

import "sort"

func hIndex(citations []int) int {
	// Consider a list that is sorted in descending order, ex. [6, 5, 3, 1]
	// We want to find the number of values higher than their position in the array, so here
	// if the 3 was a 4 then we would have 6, 5, and 4 all higher than their positions (1, 2, and 3)
	// but if the 3 was a 2 then it would NOT be higher than its index and the answer would be 2

	// An O(nlogn) solution would be to sort it and then iterate until we find this value
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	h := 0
	for index, num := range citations {
		if num <= index {
			break
		} else {
			h++
		}
	}
	return h
}
