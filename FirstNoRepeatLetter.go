package main

import (
	"math"
	"strings"
)

// https://www.codewars.com/kata/52bc74d4ac05d0945d00054e/train/go
func FirstNonRepeating(str string) string {
	incasestr := strings.ToLower(str)
	repeatLetters := make(map[rune]struct{})
	norepeatLetters := make(map[rune]int)
	for idx, v := range incasestr {
		if _, ok := repeatLetters[v]; ok {
			continue
		} else {
			if _, ok := norepeatLetters[v]; ok {
				delete(norepeatLetters, v)
				repeatLetters[v] = struct{}{}
			} else {
				norepeatLetters[v] = idx
			}
		}
	}

	norepeatLettersLen := len(norepeatLetters)
	if norepeatLettersLen > 0 {
		// 找到非重复字符里面索引最小的那个
		min := math.MaxInt32
		for _, idx := range norepeatLetters {
			if idx < min {
				min = idx
			}
		}
		return string(str[min])
	}

	return ""
}

// 最简短的，~,~
// func FirstNonRepeating(str string) string {
//     for _, c := range str {
//         if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
//             return string(c)
//         }
//     }
//     return ""
// }
