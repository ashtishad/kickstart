package main

import "fmt"

func findRem(k int, candies []int) int {
	var c int // total candies
	for _, v := range candies {
		c += v
	}
	if c%k == 0 {
		return 0
	}

	return c % k
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		var N, M int
		fmt.Scanf("%d %d", &N, &M)
		var C []int
		for j := 0; j < N; j++ {
			var c int
			fmt.Scanf("%d", &c)
			C = append(C, c)
		}
		fmt.Printf("Case #%d: %d\n", i+1, findRem(M, C))
	}
}
