package webart

import (
	"fmt"
	"strings"
)

// CheckAscii takes word a string and returns
// true if the characters are printable
func CheckAscii(word string) bool {
	word = strings.ReplaceAll(word, "\r\n", " ")
	wordsArr := strings.Split(word, "")
	for _, v := range wordsArr {
		for _, val := range v {
			if val < 32 || val > 126 {
				fmt.Println("found a non-ascii character")
				return false
			}
		}
	}
	return true
}
