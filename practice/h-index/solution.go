package main

import "fmt"

func getHIndex(citations []int) int {
	n := len(citations)

	buckets := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}

	var h int
	for i := n; i >= 0; i-- {
		h += buckets[i]
		if h >= i {
			return i
		}
	}
	return 0
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		citations := make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &citations[j])
		}

		fmt.Printf("Case #%d:", i)
		for j := 0; j < n; j++ {
			fmt.Printf(" %d", getHIndex(citations[:j+1]))
		}
		fmt.Println()
	}
}
