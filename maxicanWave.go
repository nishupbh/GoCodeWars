package main

import (
	"fmt"
	"strings"
)

func wave(words string) []string {

	var mexWave []string

	if len(words) == 0 {
		return nil
	} else {
		//mexWave = append(mexWave)
		for i := 0; i < len(words); i++ {
			var tempWord []string
			if words[i] == 32 {
				continue
			}
			for j := 0; j < len(words); j++ {
				if (j == i) && (words[i] > 96 && words[i] < 123) {
					tempWord = append(tempWord, string(words[j]-32))
				} else {
					tempWord = append(tempWord, string(words[j]))
				}
			}
			//fmt.Println(tempWord)
			mexWave = append(mexWave, strings.Join(tempWord, ""))
		}

		return mexWave
	}
}

func main() {
	fmt.Println(wave(""))
}
