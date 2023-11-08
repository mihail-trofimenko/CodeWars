package main

import (
	"fmt"
	"strconv"
)

func getBits(b byte) [8]byte {
	var sb [8]byte
	for i := 0; i < 8; i++ {
		sb[i] = byte(int(b)>>i) & 1
	}
	return [8]byte(sb)
}

func inject(orig [3][3]byte, mask [8]byte) [3][3]byte {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i+j < 8 {
				orig[i][j] = orig[i][j] | mask[i+j]
			}
		}
	}
	return orig
}

func Conceal(msg string, pixels [][]uint8) [][]uint8 {

	// your code here
	return nil
}

type pixel3 [3][3]byte

func main() {

	var n byte = 204
	var sn [8]byte
	var tp = pixel3{{254, 254, 254}, {254, 254, 254}, {254, 254, 254}}
	fmt.Println(strconv.FormatInt(int64(n), 2), " - MASK")

	sn = getBits(n)
	for i := 0; i < 8; i++ {
		fmt.Print(strconv.FormatInt(int64(sn[i]), 2))
	}

	fmt.Println(" - MASK in array")

	tp = inject(tp, sn)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(strconv.FormatInt(int64(tp[i][j]), 2))
		}

	}

}
