package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prod := make([]int, n)
	prod[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		prod[i] = prod[i+1] * nums[i+1]
	}

	acc := 1
	for i := 1; i < n; i++ {
		acc = acc * nums[i-1]
		prod[i] = prod[i] * acc
	}
	return prod
}
