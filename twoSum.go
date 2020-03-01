package main

import (
	"fmt"
	"io"
)

//TwoSum will fetch the index of matching number
func TwoSum(number []int, target int) [2]int {
	for i := 0; i < len(number)-1; i++ {
		for j := i + 1; j < len(number); j++ {
			if number[i]+number[j] == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
	io.Copy()

}

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3, 4}, 4))
}
