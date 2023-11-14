package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SpinWords(str string) string {
	words := strings.Fields(str)
	str = ""
	for idx, word := range words {
		if len(word) > 4 {
			words[idx] = Reverse(word)
		}
		str += words[idx]
		if idx < len(words)-1 {
			str += " "
		}
	}

	return str
} // SpinWords

func main() {
	refString := "Burgers are my favorite fruit"

	fmt.Println(SpinWords(refString))
}
