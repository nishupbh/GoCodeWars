package main

import (
	"fmt"
	"strings"
)

func High(s string) string {
	strSlice := strings.Split(strings.ToLower(s), " ")
	var highScoringWord string
	var highScore, temp int32
	for _, value := range strSlice {
		temp = 0
		for _, strValue := range value {
			temp += (strValue - 96)
			//fmt.Println(strValue - 96)
		}
		if highScore < temp {
			highScore = temp
			highScoringWord = value

			//fmt.Printf("High Score %v with  word %v \n", highScore, highScoringWord)
		}
	}
	return highScoringWord
}

func main() {

	fmt.Println(High("wiwwjvhcxf qhwpstkqv"))
}
