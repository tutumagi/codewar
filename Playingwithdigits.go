package main

import "math"

func DigPow(n, p int) int {
	factor := make([]int, 0, 10)
	tmp := n
	for tmp > 0 {
		factor = append(factor, tmp%10)
		tmp /= 10
	}
	msum := 0
	for i := len(factor) - 1; i >= 0; i-- {
		msum += int(math.Pow(float64(factor[i]), float64(p)))
		p++
	}
	if msum%n == 0 {
		return msum / n
	}

	return -1
}
