package main

import (
	"fmt"
	"strconv"
)

func MultiplyDigits(n int) int {
	res := 1
	numStr := strconv.Itoa(n)
	for _, r := range numStr {
		num, _ := strconv.Atoi(string(r))
		res *= num
	}
	return res
}

func Persistence(n int) int {
	i := 0
	for i = 0; n > 9; i++ {
		n = MultiplyDigits(n)
	}
	return i
}

func main() {
	fmt.Println(Persistence(39))
}
