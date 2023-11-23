package main

import "fmt"

func FindOdd(seq []int) int {

	for i := 0; i < len(seq); i++ {
		oddcheck := 1
		for j := 0; j < len(seq); j++ {
			if seq[i] == seq[j] && i != j {
				oddcheck++
			}
		}
		if oddcheck == 1 || oddcheck%2 == 1 {
			return seq[i]
		}
	}

	return 0
}

func main() {
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
}
