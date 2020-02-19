package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	phoneList := make([]string, 14)
	phoneList = append(phoneList, "(")
	for i := 0; i < 10; i++ {
		phoneList = append(phoneList, strconv.FormatUint(uint64(numbers[i]), 10))
		if i == 2 {
			phoneList = append(phoneList, ") ")
		}
		if i == 5 {
			phoneList = append(phoneList, "-")
		}

	}
	return strings.Join(phoneList, "")

}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
