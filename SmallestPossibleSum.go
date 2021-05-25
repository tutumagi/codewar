package main

// https://www.codewars.com/kata/52f677797c461daaf7000740/train/go
func SmallestPossibleSum(ar []int) int {
	length := len(ar)
	// 求最大公约数，然后 * len
	result := ar[length-1]
	for i := length - 2; i >= 0; i-- {
		result = gcd(result, ar[i])
	}

	return result * length
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type SortByInt []int

func (a SortByInt) Len() int           { return len(a) }
func (a SortByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByInt) Less(i, j int) bool { return a[i] < a[j] }
