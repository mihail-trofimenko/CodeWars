package main

import (
	"fmt"
	"strings"
)

func Rot13(msg string) string {

	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}

	return strings.Map(rot13, msg)
}

func main() {
	fmt.Println(Rot13("'Twas brillig and the slithy gopher..."))
}
