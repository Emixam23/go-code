package top_k

func topKFrequent(nums []int, k int) []int {
	// 1. Create a frequency map to count occurrences of each number
	var buckets = make(map[int]int)
	for _, n := range nums {
		buckets[n]++
	}

	// 2. Create a slice of slices to hold numbers by their frequency
	var sBuckets = make([][]int, len(nums)+1)
	for n, freq := range buckets {
		sBuckets[freq] = append(sBuckets[freq], n)
	}

	// 3. Collect the top k frequent elements
	var res = make([]int, 0, k)
	for i := len(sBuckets) - 1; i >= 0 && len(res) < k; i-- {
		// 4. For the same frequency, add all numbers to the result
		for _, n := range sBuckets[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
