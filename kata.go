package main

import "fmt"

func Solution(str string) []string {

	var str1 []string

	if len(str)%2 == 0 {
		for i := 0; i < len(str); i = i + 2 {
			str1 = append(str1, str[i:i+2])
		}
		return str1
	} else {
		for i := 0; i < len(str)-1; i = i + 2 {
			str1 = append(str1, str[i:i+2])
		}
		str1 = append(str1, string(str[len(str)-1])+"_")
		return str1
	}
}

func main() {
	fmt.Println(Solution("abcdefghjh"))
}
