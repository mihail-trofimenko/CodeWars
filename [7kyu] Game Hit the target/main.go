package main

import "fmt"

var arrow struct {
	x     int
	y     int
	exist bool
}

func Solution(mtrx [][]rune) bool {

	for y := 0; y < len(mtrx); y++ {
		for x := 0; x < len(mtrx[y]); x++ {
			if mtrx[y][x] == '>' {
				arrow.y = y
				arrow.x = x
				arrow.exist = true
			}
		}
	}
	if arrow.exist {
		for x := arrow.x; x < len(mtrx[arrow.y]); x++ {
			if mtrx[arrow.y][x] == 'x' {
				return true
			}
		}
	}
	return false
}

func main() {

	mtrx := [][]rune{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{'>', ' ', ' '},
		{' ', ' ', ' '}}

	fmt.Println(Solution(mtrx))

}
