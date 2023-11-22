package main

import "fmt"

func Snail(snaipMap [][]int) (result []int) {
	// start direction and position
	dx, dy, x, y := 1, 0, -1, 0
	// margins of non-scanned area
	var my0, my1, mx0, mx1 int
	my1 = len(snaipMap) - 1
	mx1 = len(snaipMap) - 1

	if len(snaipMap[0]) == 0 { // if empty matrix
		result = make([]int, 0)
		return result
	}

	for i := 0; i < (len(snaipMap) * len(snaipMap)); i++ {
		x += dx
		y += dy
		result = append(result, snaipMap[y][x])

		// moving right - check margin, change direction and move margin
		if dx == 1 && x+dx == mx1+1 {
			dx = 0
			dy = 1
			my0++
			continue
		}
		// moving left - check margin, change direction and move margin
		if dx == -1 && x+dx == mx0-1 {
			dx = 0
			dy = -1
			my1--
			continue
		}
		// moving down - check margin, change direction and move margin
		if dy == 1 && y+dy == my1+1 {
			dx = -1
			dy = 0
			mx1--
			continue
		}
		// moving up - check margin, change direction and move margin
		if dy == -1 && y+dy == my0-1 {
			dx = 1
			dy = 0
			mx0++
			continue
		}

	}
	return result
}

func main() {
	var testArray1 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	fmt.Println(Snail(testArray1))
}
