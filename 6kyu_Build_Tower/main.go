package main

import "fmt"

func TowerBuilder(nFloors int) (pyramid []string) {

	baseWidth := (nFloors-1)*2 + 1
	center := nFloors - 1
	var blocks []rune

	for i := 0; i < baseWidth; i++ {
		blocks = append(blocks, ' ')
	}

	for i := 0; i < center+1; i++ {
		blocks[center+i] = '*'
		blocks[center-i] = '*'
		pyramid = append(pyramid, string(blocks))
	}

	return pyramid
}

func main() {
	pyramid := TowerBuilder(50)
	for {
		for _, r := range pyramid {
			fmt.Println(r)
		}
		for i := len(pyramid) - 1; i >= 0; i-- {
			fmt.Println(pyramid[i])
		}
	}
}
