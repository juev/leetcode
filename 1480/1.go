func runningSum(nums []int) []int {
	var result []int
	prev := int(0)
	for _, val := range nums {
		result = append(result, val+prev)
		prev = prev + val
	}
	return result
}