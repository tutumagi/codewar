package main

// https://www.codewars.com/kata/51e0007c1f9378fa810002a9/go
func MakeTheDeadfishSwim(data string) []int {
	ret := make([]int, 0)
	type cmd func(num *int)
	cmds := map[rune]cmd{
		'i': func(num *int) { *num++ },
		'd': func(num *int) { *num-- },
		's': func(num *int) { *num = (*num) * (*num) },
		'o': func(num *int) { ret = append(ret, *num) },
	}
	num := 0
	for _, r := range data {
		if c, ok := cmds[r]; ok {
			c(&num)
		}
	}
	return ret
}
