package main

import (
	"math"
	"strings"
)

// https://www.codewars.com/kata/5508249a98b3234f420000fb
func MovingShift(s string, shift int) []string {
	per := int(math.Ceil(float64(len(s)) / float64(5)))
	// rest := shift - 4*per
	shiftinit := shift
	result := make([]string, 5)
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			// 算出 ['a', 'z'] 范围的字符
			r = rune(((int(r)+shift)-'a')%26 + int('a'))
		} else if r >= 'A' && r <= 'Z' {
			// 算出 ['A', 'Z'] 范围的字符
			r = rune(((int(r)+shift)-'A')%26 + int('A'))
		}
		ii := (shift - shiftinit) / per
		result[ii] = strings.Join([]string{result[ii], string(r)}, "")
		shift++
	}

	return result
}
func DemovingShift(arr []string, shift int) string {
	result := ""
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			r := arr[i][j]
			if r >= 'a' && r <= 'z' {
				absBase := (int(r) - shift - 'a')
				r = byte((absBase+(int(math.Abs(float64(absBase)))/26+1)*26)%26 + 'a')
				// 算出 ['a', 'z'] 范围的字符
				// r = byte((int(r)-shift-'a'+26)%26 + 'a')
				// r = byte((int(r)-shift-'a'+26)%26 + 'a')

			} else if r >= 'A' && r <= 'Z' {
				// 算出 ['A', 'Z'] 范围的字符
				// r = (r-'A')%26 + 'A' - byte(shift)
				absBase := (int(r) - shift - 'A')
				r = byte((absBase+(int(math.Abs(float64(absBase)))/26+1)*26)%26 + 'A')

			}
			result += string(r)
			shift++
		}
	}

	return result
}
