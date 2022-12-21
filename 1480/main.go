func runningSum(nums []int) []int {
	prev := int(0)
	for i, val := range nums {
		prev = prev + val
		nums[i] = prev

	}
	return nums
}
