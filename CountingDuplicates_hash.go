package main

import (
	"fmt"
	"strings"
)

func duplicate_counts(s string) (c int) {
	h := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return
}

func main() {
	fmt.Println(duplicate_counts("Indivisibilities"))
}
