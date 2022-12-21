func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]uint, 0)
	for _, ch := range magazine {
		m[ch]++
	}
	for _, ch := range ransomNote {
		if m[ch] == 0 {
			return false
		}
		m[ch]--
	}
	return true
}