package main

import (
	"math/big"
)

// 		1 2 3 4 5 6 7 8 9 0
//0		0 0 0 0 0 0 0 0 0 1
//1		1 1 1 1 1 1 1 1 1 1
//2 	2 4 8 6 2 4 8 6 2 1
//3		3 9 7 1 3 9 7 1 3 1
//4 	4 6 4 6 4 6 4 6 4 1
//5		5 5 5 5 5 5 5 5 5 1
//6		6 6 6 6 6 6 6 6 6 1
//7		7 9 3 1 7 9 3 1 7 1
//8 	8 4 2 6 8 4 2 6 8 1
//9		9 1 9 1 9 1 9 1 9 1
// 底数：尾数顺序
var predefine = map[int][]int{
	0: {0},
	1: {1},
	2: {2, 4, 8, 6},
	3: {3, 9, 7, 1},
	4: {4, 6},
	5: {5},
	6: {6},
	7: {7, 9, 3, 1},
	8: {8, 4, 2, 6},
	9: {9, 1},
}

// https://www.codewars.com/kata/5518a860a73e708c0a000027/train/go
// https://brilliant.org/wiki/finding-the-last-digit-of-a-power/
// https://www.youtube.com/watch?v=k7rm55Sw-SE
func LastDigitWithHugeNumber(as []int) int {
	if len(as) == 0 {
		return 1
	}
	result := big.NewInt(1)
	four := big.NewInt(4)
	t1 := new(big.Int)
	tmp := new(big.Int)
	for i := len(as) - 1; i >= 0; i-- {
		if result.Cmp(four) >= 0 {
			result = result.Mod(result, four)
			result = result.Add(result, four)
		}
		t1.SetInt64(int64(as[i]))
		result = tmp.Exp(t1, result, nil)
	}

	return int(result.Mod(result, big.NewInt(10)).Int64())
}
