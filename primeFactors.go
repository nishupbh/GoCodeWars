package main

import (
	"fmt"
	"strconv"
	"time"
)

//func PrimeFactors(n int) string {
//	var str string

//	for i := 2; i <= n; i++ {
//		flag := 1
//		for j := 2; j <= i/2; j++ {
//			if i%j == 0 {
//				flag = 0
//				break
//			}
//		}
//		if flag == 1 {
//			count := 0
//			for n%i == 0 {
//				count++
//				n /= i
//			}
//			if count == 1 {
//				str += "(" + strconv.Itoa(i) + ")"
//			} else if count > 1 {
//				str += "(" + strconv.Itoa(i) + "**" + strconv.Itoa(count) + ")"
//			}
//		}
//	}
//	return str
//}

func PrimeFactors(n int) string {
	var str string
	for i := 2; i <= n; i++ {
		count := 0
		for n%i == 0 {
			count++
			n /= i
		}
		if count == 1 {
			str += "(" + strconv.Itoa(i) + ")"
		} else if count > 1 {
			str += "(" + strconv.Itoa(i) + "**" + strconv.Itoa(count) + ")"
		}
	}
	return str
}

func main() {
	start := time.Now()
	fmt.Println(PrimeFactors(7775460))
	fmt.Println("Time taken:", time.Since(start))
}
