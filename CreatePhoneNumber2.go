package main

import "fmt"

func CreatePhoneNumber1(n [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
}

func main() {
	fmt.Println(CreatePhoneNumber1([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
