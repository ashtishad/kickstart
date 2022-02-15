package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for t := 1; t <= T; t++ {
		var k string
		fmt.Scanf("%s\n", &k)

		fmt.Printf("Case #%d: %s %s %s%s\n", t, k, "is ruled by", getRuler(k), ".")
	}
}

func getRuler(k string) string {
	last := k[len(k)-1]

	switch last {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return "Alice"
	case 'y', 'Y':
		return "nobody"
	default:
		return "Bob"
	}
}
