package main

import (
	"fmt"
	"strconv"
)

func DigitalRoot(n int) (sum int) {
	numstr := strconv.Itoa(n)
	for {
		sum = 0
		for i := 0; i < len(numstr); i++ {
			sum += int(numstr[i] - 48)
		}
		if sum < 10 {
			return sum
		}
		numstr = strconv.Itoa(sum)
	}
}

func main() {
	fmt.Println(DigitalRoot(167346))

}
