package twoSum

func hasDuplicate(nums []int) bool {
	var m = map[int]bool{}

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}

	return false
}
