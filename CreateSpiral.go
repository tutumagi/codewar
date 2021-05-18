package main

// https://www.codewars.com/kata/536a155256eb459b8700077e
func CreateSpiral(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		if len(result[i]) == 0 {
			result[i] = make([]int, n)
		}
	}

	lastDir := 1 // 1 right, 2 down, 3 left, 4 top
	row := 0
	col := 0
	canright := func() bool { return col+1 < n && result[row][col+1] == 0 }
	candown := func() bool { return row+1 < n && result[row+1][col] == 0 }
	canleft := func() bool { return col-1 >= 0 && result[row][col-1] == 0 }
	cantop := func() bool { return row-1 < n && result[row-1][col] == 0 }

	for i := 1; i <= n*n; i++ {
		result[row][col] = i

		if lastDir == 1 && canright() {
			col++
		} else if lastDir == 2 && candown() {
			row++
		} else if lastDir == 3 && canleft() {
			col--
		} else if lastDir == 4 && cantop() {
			row--
		} else {
			if canright() { // 尝试走右方向
				col++
				lastDir = 1
			} else if candown() { // 尝试走下方向
				row++
				lastDir = 2
			} else if canleft() { // 尝试走左方向
				col--
				lastDir = 3
			} else { // 尝试走上方向
				row--
				lastDir = 4
			}
		}

	}
	return result
}
