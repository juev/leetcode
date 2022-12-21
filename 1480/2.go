func runningSum(nums []int) []int {
	prev := int(0)
	for i, val := range nums {
		nums[i] = nums[i] + prev
		prev = prev + val
	}
	return nums
}