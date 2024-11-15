package webart

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

// Filereader takes filename a string and returns the content as a slice []string
// if an error is encounterd. It is also logged
func FileReader(filename string) ([]string, error) {
	var content []string
	stdCheckSum := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shdCheckSum := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thkCheckSum := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"

	file, err := os.ReadFile("./resources/" + filename + ".txt")
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write(file)
	sum := h.Sum(nil)
	checkSum := fmt.Sprintf("%x", sum)

	if checkSum != stdCheckSum && checkSum != shdCheckSum && checkSum != thkCheckSum {
		return nil, fmt.Errorf(" Invalid file or modified file")
	}

	if checkSum == thkCheckSum {
		content = strings.Split(string(file), "\r\n")
	} else {
		content = strings.Split(string(file), "\n")
	}

	return content, nil
}
