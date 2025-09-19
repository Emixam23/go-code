package encode_and_decode

import (
	"fmt"
	"strconv"
)

type Solution struct {
}

func (s *Solution) Encode(strs []string) string {
	var res string
	for _, str := range strs {
		res += fmt.Sprintf("%d#%s", len(str), str)
	}
	return res
}

func (s *Solution) Decode(str string) []string {
	var res []string
	var i, j int

	for i < len(str) {
		// find the length of the string
		for i < len(str) && str[i] != '#' {
			currentChar := string(rune(str[i]))
			_ = currentChar
			i++
		}
		if i == len(str) {
			break
		}
		length, _ := strconv.Atoi(str[j:i])

		// move to the start of the string
		i++
		j = i + length
		res = append(res, str[i:j])

		// move to the next length
		i = j
	}
	return res
}

// 5#hello5#world
