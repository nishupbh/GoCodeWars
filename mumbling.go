package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	var list []string
	//var temp string
	s = strings.ToLower(s)
	splitString := strings.Split(s, "")
	for i, value := range splitString {
		//temp = strings.ToUpper(value) + strings.Repeat(value,i)

		list = append(list, strings.ToUpper(value)+strings.Repeat(value, i))
	}

	s = strings.Join(list, "-")
	return s
}

func main() {
	fmt.Println(Accum("abcd"))
}
