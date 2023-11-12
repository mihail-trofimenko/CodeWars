package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FindTheNumberPlate(n int) (plate string) {

	plate = ""

	mul := 26 * 26 * 999
	RuneSeries := int(n / (mul))
	plate += string('a' + rune(RuneSeries))
	if n >= RuneSeries*mul {
		n -= RuneSeries * mul
	}

	mul = 26 * 999
	RuneSeries = int(n / mul)
	plate += string('a' + rune(RuneSeries))
	if n >= RuneSeries*mul {
		n -= RuneSeries * mul
	}
	RuneSeries = int(n / (999))
	plate += string('a' + rune(RuneSeries))
	if n >= RuneSeries*999 {
		n -= RuneSeries * 999
	}
	plate = Reverse(plate)
	plate += fmt.Sprintf("%03d", n+1)
	return plate
}

func main() {

	fmt.Println(FindTheNumberPlate(17558423))
}
