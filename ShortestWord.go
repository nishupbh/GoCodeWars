package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	sliceStr := strings.Split(s, " ")
	tempLen := len(sliceStr[0])
	for _, value := range sliceStr {
		if tempLen > len(value) {
			tempLen = len(value)
		}
	}
	return tempLen
}

func main() {
	fmt.Println("Shortest word length is ", FindShort("String will never be empty and you do not need to account for different data types"))
}
