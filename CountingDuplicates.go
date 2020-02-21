package main

import (
	"fmt"
	"strings"
)

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func duplicate_count(s1 string) int {
	max_val := 0
	fmt.Println(unique(strings.Split(strings.ToLower(s1), "")))
	for _, value := range unique(strings.Split(strings.ToLower(s1), "")) {
		count := 0
		for _, value1 := range strings.Split(strings.ToLower(s1), "") {

			if value == value1 {
				count++
			}
		}
		if count > 1 {
			max_val++
		}
	}
	return max_val
}

func main() {
	fmt.Println(duplicate_count("Indivisibilities"))
}
