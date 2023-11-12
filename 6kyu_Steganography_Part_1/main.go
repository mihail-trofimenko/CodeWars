package main

import (
	"fmt"
	//"strconv"
)

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

// пихаем по байту входной строки в каждые 3 пикселя
func Conceal(msg string, pixels [][]uint8) [][]uint8 {

	var tp = [][]uint8{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}} // массив из трёх пикселей RGB
	var mask [8]byte                                    // маска для наложения на пиксель

	if int((len(pixels))/3) < len(msg) {
		return nil
	}

	for i := 0; i < len(msg); i++ {
		for b := 0; b < 8; b++ { // получаем маску из очередного символа строки
			mask[7-b] = byte(int(msg[i])>>b) & 1
		}

		for j := 0; j < 3; j++ { // всасываем 3 пикселя из массива
			tp[j] = pixels[i*3+j]
		}

		tp = inject(tp, mask) // накладываем маску на пиксели

		for j := 0; j < 3; j++ {
			pixels[i*3+j] = tp[j] // кладём пиксели взад
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
