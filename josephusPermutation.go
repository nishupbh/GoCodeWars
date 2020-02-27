package main

import "fmt"

func Josephus(items []interface{}, k int) []interface{} {
	executionOrder := []interface{}{}
	length := len(items)
	for true {
		if k > len(items) {
			for _, val := range items {
				executionOrder = append(executionOrder, val)
			}
			break
		} else {
			if k == len(items) {
				executionOrder = append(executionOrder, items[k-1])
				items = items[:k]
			} else {
				i := k
				executionOrder = append(executionOrder, items[i-1])
				i = i + k - 1
				items = append(items[:i-1], items[i:])
				if i > len(items) {
					i = i % length
				}
			}
		}
	}
	return executionOrder

}
func main() {
	//items := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7}, 3))
}
