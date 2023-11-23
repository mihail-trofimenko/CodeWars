package main

import "fmt"

func Solution(list []int) (res string) {
	var rangeStart, rangeEnd int // This variables here is for better understanding of my logic
	fmt.Println(list)
	for i := 0; i < len(list); i++ {
		if i < len(list)-2 && list[i+1] == list[i]+1 && list[i+2] == list[i]+2 {
			rangeStart = list[i]
			rangeEnd = 0
			res += fmt.Sprintf("%d-", rangeStart)
			for j := i + 1; j < len(list); j++ {
				if j == len(list)-1 && list[j] == (list[j-1]+1) {
					res += fmt.Sprintf("%d", list[j])
					return res
				}
				if j < len(list) && list[j] != (list[j-1]+1) {
					i = j - 1
					rangeEnd = list[i]
					break
				}
			}
			res += fmt.Sprintf("%d,", rangeEnd)
			continue
		}
		res += fmt.Sprintf("%d,", list[i])
	}
	return res[:len(res)-1]
}

func main() {
	//fmt.Print(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
	//-6,-3-1,3-5,7-11,14,15,17-20
	test1 := []int{58, 60, 61, 62, 63, 64, 66, 67, 68, 75, 76, 77, 83}
	fmt.Print(Solution(test1))
}
