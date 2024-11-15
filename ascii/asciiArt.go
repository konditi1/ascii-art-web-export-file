package webart

import (
	"fmt"
	"strings"
)

// Ascii takes fileArr a slice of string and words a string
// Ascii returns a string. The string is words in Ascii-art
func Ascii(fileArr []string, words string) (string, error) {
	var result string

	if words == "" {
		return "", fmt.Errorf("Error")
	}
	words = strings.ReplaceAll(words, "\r\n", "\n")
	wordsArr := strings.Split(words, "\n")
	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					result += fileArr[int(start)+i]
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result, nil
}
