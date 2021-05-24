package main

// https://www.codewars.com/kata/546d5028ddbcbd4b8d001254/
// https://www.codewars.com/kata/explosive-sum/
// 1	[1]		1
// 2    [2] [1, 1]	2
// 3    [3] [2, 1] [1,1,1] 3
// 4    [4] [3,1] [2,2] [2,1,1] [1,1,1,1]  5
// 5    [5] [4, 1] [3, 2] [3,1,1] [2,2,1][2,1,1,1][1,1,1,1,1] 7
// https://en.wikipedia.org/wiki/Partition_(number_theory)#
// NOTE: the answer just from wiki.
const maxLength = 2000

var partitionsCache = make([]int, maxLength+1)

func init() {
	partitionsCache[0] = 1
	for i := 1; i <= maxLength; i++ {
		for j := i; j <= maxLength; j++ {
			partitionsCache[j] += partitionsCache[j-i]
		}
	}
}
func Partitions(n int) int {
	return partitionsCache[n]
}
