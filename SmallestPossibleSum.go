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
