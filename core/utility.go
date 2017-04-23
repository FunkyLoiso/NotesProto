package core

import (
	"bufio"
	"bytes"
	"log"
	"strings"
)

const (
	titleLength = 30
)

func MakeTitle(text string) string {
	log.Println("MakeTitle for", text)
	// remover := NewMarkdownRemover(text)
	remover := strings.NewReader(text)
	lineScanner := bufio.NewScanner(remover)
	log.Println("lineScanner is here")
	lineScanner.Scan()
	if err := lineScanner.Err(); err != nil {
		log.Println("Failed to read line:", err)
		return ""
	}
	log.Println("Line:", lineScanner.Text())
	// now we have a line
	wordScanner := bufio.NewScanner(bytes.NewReader(lineScanner.Bytes()))
	wordScanner.Split(bufio.ScanWords)

	var result string
	for wordScanner.Scan() {
		word := wordScanner.Text()
		log.Println("Word:", word)
		if len(result)+len(word)+1 <= titleLength {
			result += " " + word
		} else {
			break // we've accumulated maximum possible title
		}
	}

	if err := wordScanner.Err(); err != nil {
		log.Println("Failed to read word:", err)
		return ""
	}
	log.Println("Title is", result)
	return result
}
