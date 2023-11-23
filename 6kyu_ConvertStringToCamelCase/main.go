package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func ToCamelCase(s string) string {

	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	words := strings.Fields(s)
	for i := 1; i < len(words); i++ {
		r := []rune(words[i])
		r[0] = unicode.ToUpper(r[0])
		words[i] = string(r)
	}
	s = ""
	for _, w := range words {
		s += w
	}
	return s
}

func main() {
	fmt.Print("convertStringToCamelCase >")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(ToCamelCase(string(scanner.Text())))
	fmt.Println("Now copy result and press any key to close terminal...")
	scanner.Scan()
}
