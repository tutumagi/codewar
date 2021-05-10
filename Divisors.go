package main

// func Divisors(n int) int {
// 	max := n / 2
// 	fmt.Printf("计算所有可以整除 %d 的数的数目\n", n)
// 	counter := 0
// 	mark := make([]bool, n+1)
// 	markFn := func(i int) {
// 		if !mark[i] {
// 			fmt.Printf("%d 可以整除 %d\n", i, n)
// 			mark[i] = true
// 			counter++
// 		}
// 	}
// 	// 1 和 他自身肯定可以整除的
// 	markFn(1)
// 	markFn(n)
// 	for i := 2; i <= max; i++ {
// 		if mark[i] {
// 			continue
// 		}
// 		// 如果可以整除
// 		if n%i == 0 {
// 			markFn(i)
// 			j := 2
// 			divise := i * j
// 			// 所有除数 n * i, i 属于 [1,n] 之间的数就都可以整除
// 			for ; divise < max; j++ {
// 				divise = i * j
// 				markFn(divise)
// 			}
// 		}
// 	}
// 	return counter
// }

func Divisors(n int) int {
	if n == 1 {
		return 1
	}
	counter := 2
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			counter++
		}
	}
	return counter
}
