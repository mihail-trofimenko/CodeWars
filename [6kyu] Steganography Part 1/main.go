package main

import (
	"fmt"
	"strconv"
)

// заполняем массив, содержащий маску соответственно старшинству разрядов в исходном байте:
// номер элемента массива - это степень 2 (базы двоичной системы)

func getBits(b byte) [8]byte {
	var sb [8]byte
	for i := 0; i < 8; i++ {
		sb[7-i] = byte(int(b)>>i) & 1
	}
	return [8]byte(sb)
}

func inject(orig [3][3]byte, mask [8]byte) [3][3]byte {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				break
			} else {
				orig[i][j] = orig[i][j] | mask[i*3+j]
			}
		}
	}
	return orig
}

type pixel3 [3][3]byte

func PrintPixel(tp pixel3) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				break
			} else {
				fmt.Printf("[%d][%d]:%s\t", i, j, strconv.FormatInt(int64(tp[i][j]), 2))
			}
		}
		fmt.Println()
	}
}

// пихаем по байту в каждые 3 пикселя
func Conceal(msg string, pixels [][]uint8) [][]uint8 {

	// из условия задачи:
	// Your code should return nil/None if there
	// aren't enough pixels to conceal the message.
	if int((len(pixels))/3) < len(msg) {
		return nil
	}
	/*
		var n byte = 204 // байт для инжекции
		var sn [8]byte   // маска из 8 байтов, полученная из n
		var tp pixel3    // входной массив, трёх пикселей RGB
	*/

}

func main() {

	var n byte = 204 // байт для инжекции
	var sn [8]byte   // маска из 8 байтов, полученная из n

	// входной массив, трёх пикселей RGB
	var tp = pixel3{{254, 254, 254}, {254, 254, 254}, {254, 254, 254}}

	fmt.Println(strconv.FormatInt(int64(n), 2), " - MASK")

	sn = getBits(n)
	for i := 0; i < 8; i++ {
		fmt.Print(strconv.FormatInt(int64(sn[i]), 2))
	}

	fmt.Println(" - MASK in array")

	fmt.Println("Исходный массив:")
	PrintPixel(tp)

	tp = inject(tp, sn)

	fmt.Println("Массив под маской:")
	PrintPixel(tp)
}
