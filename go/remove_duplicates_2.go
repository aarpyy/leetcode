package leetcode

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	// left is the pointer to the position in the array where we are going to place a new element IF it passes
	left := 1

	// count is the number of times we have seen this
	count := 1

	// Our right pointer is the element we are currently considering keeping
	for right := 1; right < n; right++ {
		// If we are seeing a duplicate, check if it's too many
		if nums[left-1] == nums[right] {
			count++
			// If okay number, set value and increment both
			if count <= 2 {
				nums[left] = nums[right]
				left++
			}
		} else {
			count = 1
			nums[left] = nums[right]
			left++
		}
	}

	return left
}
