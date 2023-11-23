package main

import "fmt"

func FindOutlier(integers []int) int {

	outlierSign := 0
	oddVote := 0
	evenVote := 0

	for i := 0; i < 3; i++ {

		if integers[i]%2 == 1 || integers[i]%2 == -1 {
			oddVote++
		} else {
			evenVote++
		}
	}

	if oddVote < evenVote {
		outlierSign++
	}

	for _, r := range integers {
		if r%2 == outlierSign || r%2 == outlierSign*-1 {
			return r
		}
	}
	return 0
}

func main() {
	fmt.Println(FindOutlier([]int{2, 6, 8, -10, 3}))
}
