package main

import "fmt"

func MaximumSubarraySum(numbers []int) int {
	var MaxSoFar = 0
	var max_ending_here = 0
	for i := 0; i < len(numbers); i++ {
		max_ending_here = max_ending_here + numbers[i]
		if MaxSoFar < max_ending_here {
			MaxSoFar = max_ending_here
		}

		if max_ending_here < 0 {
			max_ending_here = 0
		}
	}
	return MaxSoFar
}

func main() {

	fmt.Println(MaximumSubarraySum([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
}
