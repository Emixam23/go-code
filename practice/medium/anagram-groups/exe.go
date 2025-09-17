package is_anagram

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	// Crate a map to hold buckets of anagrams
	var buckets = make(map[string][]string)

	// Iterate over each string in the input slice
	// Sort the characters in the string to create a key (unique representation of the anagram group)
	// Append the original string to the corresponding bucket
	for _, str := range strs {
		var runes = []rune(str)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		buckets[string(runes)] = append(buckets[string(runes)], str)
	}

	// Convert the map of buckets to a slice of slices for the result
	var result = make([][]string, 0, len(buckets))
	for _, bucket := range buckets {
		result = append(result, bucket)
	}

	return result
}
