package two_sum

func twoSum(nums []int, target int) []int {
	var res = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res[0] = i
			res[1] = j
			if nums[i]+nums[j] == target {
				return res
			}
		}
	}

	return []int{0, 0}
}
