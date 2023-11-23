package main

import "fmt"

func Spiralize(size int) [][]int {

	x, y := 0, 0       // current coords
	movX, movY := 1, 0 // current increments

	sp := make([][]int, size)
	for i := 0; i < len(sp); i++ { // filling array with 0
		sp[i] = make([]int, size)
		for j := 0; j < size; j++ {
			sp[i][j] = 0
		}
	}

	rightTurns := [4][4]int{ // rightTurns[i][0-1] - current movX, movY // rightTurns[i][1-2] - new movX, movY
		{1, 0, 0, 1},
		{0, 1, -1, 0},
		{-1, 0, 0, -1},
		{0, -1, 1, 0}}

	turnRight := func() bool { // change direction function based on [][]rightTurns pattern
		for i := 0; i < 4; i++ {
			if movX == rightTurns[i][0] && movY == rightTurns[i][1] {
				movX = rightTurns[i][2]
				movY = rightTurns[i][3]
				break
			}
		}
		return sp[y+movY][x+movX] != 1
	}

	theWayIsClear := func() bool { // looking for any obstacles for next move

		if movX != 0 { // if we have horisontal direction
			if x+movX == -1 || x+movX == len(sp) { // check for array margin
				return turnRight()
			}
			if x+movX == 0 || x+movX == len(sp)-1 {
				return true
			}
			if sp[y][x+(movX*2)] == 0 { // if we find 0,0 for 2 moves ahead
				return true // continue moving forward
			} else { // if the body of spiral is ahead
				return turnRight()
			}
		}

		if movY != 0 { // similar check for vertical direction
			if y+movY == -1 || y+movY == len(sp) {
				return turnRight()
			}
			if y+movY == 0 || y+movY == len(sp)-1 {
				return true
			}
			if sp[y+(movY*2)][x] == 0 {
				return true
			} else {
				return turnRight()
			}
		}
		return false
	}

	for theWayIsClear() {
		if sp[y+movY][x+movX] != 1 {
			sp[y][x] = 1
			x += movX
			y += movY
		} else {
			break
		}
	}

	return sp
}

func main() {
	spiral := Spiralize(11)
	for _, r := range spiral {
		fmt.Println(r)
	}
}
