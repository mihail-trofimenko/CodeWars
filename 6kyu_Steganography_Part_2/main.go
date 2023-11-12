package main

import (
	"fmt"
)

func getCharFrom3Pixels(pixel_x3 [][]uint8) (result uint8) {

	for i := 0; i < 8; i++ {
		result += (uint8(pixel_x3[int(i/3)][i-int(i/3)*3]) % 2) << (7 - i)
	}
	return result
}

func Extract(pixels [][]uint8) string {
	msg := ""
	pixel_x3 := [][]uint8{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	for i := 0; i < int(len(pixels)); i += 3 { // 	перебор по тройкам пикселей
		for j := 0; j < 3; j++ { // 				перебор внутри тройки пикселей
			for c := 0; c < 3; c++ { // 			перебор цветов внутри пикселя
				pixel_x3[j][c] = pixels[i+j][c]
			}
		}
		msg += string(getCharFrom3Pixels(pixel_x3))
	}
	return msg
}

func main() {

	pixelArray := [][]uint8{
		{200, 213, 128}, {8, 71, 54}, {154, 124, 135},
		{138, 175, 117}, {64, 98, 137}, {142, 191, 5},
		{120, 113, 209}, {20, 167, 193}, {38, 146, 171},
		{88, 235, 147}, {180, 221, 49}, {24, 186, 226},
		{244, 215, 195}, {114, 19, 67}, {101, 161, 57},
	}

	fmt.Println(Extract(pixelArray))

	//fmt.Println(strconv.FormatInt(int64(getCharFrom3Pixels([][]uint8{{85, 0, 0}, {0, 0, 0}, {0, 0, 0}})), 2))
}
