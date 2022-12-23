func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	input := x
	res := x % 10
	x = int(x / 10)
	for x > 0 {
		res = (res * 10) + (x % 10)
		x = int(x / 10)
	}
	return res == input
}