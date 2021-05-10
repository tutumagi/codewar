package main

func FindOutlier(integers []int) int {
	evenCounter := 0
	oddCounter := 0
	lastOdd := 0
	lastEven := 0
	for _, v := range integers {
		if v%2 == 0 {
			evenCounter += 1
			lastEven = v
		} else {
			oddCounter += 1
			lastOdd = v
		}

		if evenCounter >= 2 && oddCounter > 0 {
			return lastOdd
		}
		if oddCounter >= 2 && evenCounter > 0 {
			return lastEven
		}
	}
	panic("")
}
