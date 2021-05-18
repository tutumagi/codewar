package main

import "strings"

func duplicate_count(s1 string) int {
	count := 0
	ss := strings.ToLower(s1)
	marked := make(map[int]int, 26+10)
	for _, v := range ss {
		if cc, ok := marked[int(v)]; ok {
			marked[int(v)]++
			if cc == 1 {
				count++
			}
		} else {
			marked[int(v)] = 1
		}
	}
	return count
}
