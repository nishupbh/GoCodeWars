package main

import (
	"fmt"
	"sort"
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

func TwoToOne(s1 string, s2 string) string {

	//s1Slice, s2Slice := unique(strings.SplitN(s1, "", len(s1))), unique(strings.SplitN(s2, "", len(s2)))
	s1Slice := unique(strings.Split(s1+s2, ""))

	sort.Sort(sort.StringSlice(s1Slice))
	return strings.Join(s1Slice, "")
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere"))

}
