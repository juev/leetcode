func numberOfSteps(num int) int {
	c := 0
	for num != 0 {
		c++
		if num%2 == 0 {
			num /= 2
			continue
		}
		num--
	}
	return c
}