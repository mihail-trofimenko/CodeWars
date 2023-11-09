package main

import (
	"fmt"
	//"strconv"
)

func getBits(b byte) [8]byte {
	var sb [8]byte
	for i := 0; i < 8; i++ {
		sb[7-i] = byte(int(b)>>i) & 1
	}
	return [8]byte(sb)
}

func inject(orig [][]uint8, mask [8]uint8) [][]uint8 {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				break
			} else {
				switch mask[i*3+j] {
				case 0:
					orig[i][j] &= 254
				case 1:
					orig[i][j] |= 1
				}
			}
		}
	}
	return orig
}

// пихаем по байту в каждые 3 пикселя
func Conceal(msg string, pixels [][]uint8) [][]uint8 {
	// массив из трёх пикселей RGB на корм inject()
	var tp = [][]uint8{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	if int((len(pixels))/3) < len(msg) {
		return nil
	}

	for i := 0; i < len(msg); i++ {
		for j := 0; j < 3; j++ {
			tp[j] = pixels[i*3+j]
		}
		tp = inject(tp, getBits(msg[i]))
		for j := 0; j < 3; j++ {
			pixels[i*3+j] = tp[j]
		}

	}
	return pixels
}

func main() {

	var pix = [][]uint8{
		{255, 0, 0}, {0, 255, 0}, {0, 0, 255},
		{255, 0, 0}, {0, 255, 0}, {0, 0, 255},
		{169, 105, 71}, {172, 211, 192}, {181, 140, 38},
		{108, 58, 63}, {105, 235, 16}, {204, 69, 21},
		{24, 40, 224}, {88, 84, 121}, {123, 41, 163},
	}
	fmt.Println(Conceal("Hello", pix))

}
