package is_anagram

func isAnagram(s string, t string) bool {
	var m = make(map[rune]int)

	// init map with chars from t
	for _, char := range t {
		m[char]++
	}

	// Subtract chars from s
	for _, char := range s {
		if count, exists := m[char]; !exists || count == 0 {
			return false
		} else if count > 0 {
			m[char]--
		}
	}

	// Check if all counts are zero
	for _, count := range m {
		if count != 0 {
			return false
		}
	}

	return true
}
