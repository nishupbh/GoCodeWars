package main

import "fmt"

func LetterChanges(str string) string {

	// code goes here
	strToByte := []byte(str)
	for i := 0; i < len(strToByte); {
		if (strToByte[i] >= 65 && strToByte[i] <= 90) || (strToByte[i] >= 97 && strToByte[i] <= 122) {
			strToByte[i]++
		}
		if strToByte[i] == 97 || strToByte[i] == 101 || strToByte[i] == 105 || strToByte[i] == 111 || strToByte[i] == 117 {
			strToByte[i] -= 32
		}
	}

	return string(strToByte)

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LetterChanges("this long cake@&"))

}
