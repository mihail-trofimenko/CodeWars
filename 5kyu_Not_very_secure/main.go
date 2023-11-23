package main

import (
	"fmt"
	"unicode"
)

func alphanumeric(str string) bool {
	if str == "" {
		return false
	}
	for _, r := range str {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("String is valid: ", alphanumeric("hello world_"))
}
