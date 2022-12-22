func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var digits []int
	for x != 0 {
		digits = append(digits, x%10)
		x = int(x / 10)
	}
	for i := 0; i < int(len(digits)/2); i++ {
		if digits[i] != digits[len(digits)-i-1] {
			return false
		}
	}
	return true
}