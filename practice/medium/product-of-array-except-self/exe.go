package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	var res = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		v := 1
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			v *= nums[j]
		}
		res[i] = v
	}

	return res
}
