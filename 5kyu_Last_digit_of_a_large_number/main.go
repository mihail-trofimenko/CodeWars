package main

import (
	"fmt"
	"math/big"
)

func LastDigit(n1, n2 string) int {
	a, b := big.NewInt(0), big.NewInt(0)
	a.SetString(n1, 10)
	b.SetString(n2, 10)

	a.Exp(a, b, big.NewInt(10))

	return int(a.Int64())
}

func main() {

	fmt.Println(LastDigit("10", "10000000000"))
}
