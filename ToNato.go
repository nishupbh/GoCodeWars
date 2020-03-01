package main

import (
	"fmt"
	"strings"
)

func ToNato(word string) string {
	var str []string
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}

	for _, val := range strings.Split(word, "") {

		for _, val1 := range nato {
			if strings.HasPrefix(val1, strings.ToUpper(val)) {
				str = append(str, val1)
			}

		}
		if val == "!" {
			str = append(str, "!")
		}
	}
	return strings.Join(str, " ")
}

func main() {
	fmt.Println(ToNato("If you can read!"))
}

/*


package kata

import "unicode"
import "strings"

func ToNato(words string) string {
  nato := []string{
    "Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot",
    "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November",
    "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor",
    "Whiskey", "Xray", "Yankee", "Zulu",
  }
  charToCharlie := map[rune]string{}
  for _, value := range nato {
    charToCharlie[rune(value[0])] = value
  }

  result := ""
  for _, letter := range words {
    if unicode.IsLetter(letter) {
      result += charToCharlie[unicode.ToUpper(letter)] + " "
    } else if unicode.IsPunct(letter) {
      result += string(letter)
    }
  }
  return strings.TrimSpace(result)
}

*/
