package main

import "fmt"

func Solution(mtrx [][]rune) bool {

	var arrow struct {
		x         int
		y         int
		dir_x     int
		direction string ""
		aimed     bool
	}

	arrow.direction = ""
	arrow.aimed = false
	fmt.Println()
	for y := 0; y < len(mtrx); y++ {
		for x := 0; x < len(mtrx[y]); x++ {
			fmt.Printf("[%s]", string(mtrx[y][x]))
		}
		fmt.Println()
	}

	for y := 0; y < len(mtrx); y++ {
		for x := 0; x < len(mtrx[y]); x++ {
			if mtrx[y][x] == '>' || mtrx[y][x] == '<' || mtrx[y][x] == 'v' || mtrx[y][x] == '^' {
				arrow.y = y
				arrow.x = x
				switch mtrx[y][x] {
				case '>':
					arrow.direction = "right"
				case '<':
					arrow.direction = "left"
				case '^':
					arrow.direction = "up"
				case 'v':
					arrow.direction = "down"
				}
				break
			}
		}
	}

	switch arrow.direction {
	case "right":
		{
			for x := arrow.x; x < len(mtrx[arrow.y]); x++ {
				if mtrx[arrow.y][x] == 'x' {
					arrow.aimed = true
				}
			}
		}

	case "left":
		{
			for x := arrow.x; x >= 0; x-- {
				if mtrx[arrow.y][x] == 'x' {
					arrow.aimed = true
				}
			}
		}

	case "down":
		{
			for y := arrow.y; y < len(mtrx); y++ {
				if mtrx[y][arrow.x] == 'x' {
					arrow.aimed = true
				}
			}
		}

	case "up":
		{
			for y := 0; y < arrow.y; y++ {
				if mtrx[y][arrow.x] == 'x' {
					arrow.aimed = true
				}
			}
		}
	}
	fmt.Println("Arrow is aimed:", arrow.aimed)
	return arrow.aimed
}

func main() {

	mtrx := [][]rune{
		{'v', ' ', ' '},
		{'x', ' ', ' '},
		{' ', ' ', ' '}}
	fmt.Println(Solution(mtrx))

	mtrx = [][]rune{
		{' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', 'v', 'x'}}
	fmt.Println(Solution(mtrx))

}
