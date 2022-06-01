func canConstruct(ransomNote string, magazine string) bool {
	hash := make(map[rune]int)

	for _, char := range magazine {
		hash[char]++
	}

	for _, char := range ransomNote {
		if hash[char] == 0 {
			return false
		} else {
			hash[char]--
		}
	}

	return true

}