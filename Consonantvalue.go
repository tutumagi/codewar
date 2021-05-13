package main

import (
	"strings"
)

// https://www.codewars.com/kata/59c633e7dcc4053512000073/go
func ConsonantValue(str string) int {
	max := 0
	tmp := 0
	for _, v := range str {
		if strings.Contains("aeiou", string(v)) {
			if tmp > max {
				max = tmp
			}
			tmp = 0
		} else {
			tmp += int(v) - 'a' + 1
		}
	}
	// we should caculate the max result again at last.
	if tmp > max {
		max = tmp
	}
	tmp = 0

	return max
}
