package main

func Multiple3And5(number int) int {
	if number < 0 {
		return 0
	}
	count3 := (number - 1) / 3
	count5 := (number - 1) / 5
	count15 := (number - 1) / 15

	sum3 := sumForMathSeq(3, count3, 3)
	sum5 := sumForMathSeq(5, count5, 5)
	sum15 := sumForMathSeq(15, count15, 15)
	return sum3 + sum5 - sum15
}

// 等差数列求和
func sumForMathSeq(from int, count int, step int) int {
	to := from + (count-1)*step
	return (from + to) * count / 2
}
