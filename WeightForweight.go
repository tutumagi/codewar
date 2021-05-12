package main

import (
	"sort"
	"strings"
)

//https://www.codewars.com/kata/55c6126177c9441a570000cc/go
type MixStr struct {
	ss     string
	weight int
}

type SortBy []*MixStr

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	left := a[i]
	right := a[j]
	if left.weight == right.weight {
		return left.ss < right.ss
	}
	return left.weight < right.weight
}

func OrderWeight(strng string) string {
	strs := strings.Split(strng, " ")
	mixs := make([]*MixStr, 0, len(strs))
	for _, s := range strs {
		ss := strings.TrimSpace(s)
		ssweight := func(sn string) int {
			w := 0
			for _, r := range sn {
				w += int(r) - '0'
			}
			return w
		}(ss)
		mixs = append(mixs, &MixStr{ss: ss, weight: ssweight})
	}

	sb := SortBy(mixs)
	sort.Sort(sb)
	result := ""
	for _, ss := range sb {
		println(ss.weight)
		result += " " + ss.ss
	}
	return strings.TrimSpace(result)
}
