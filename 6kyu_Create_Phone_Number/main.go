package main

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	return "(" + fmt.Sprint(numbers[0]) + fmt.Sprint(numbers[1]) + fmt.Sprint(numbers[2]) + ") " + fmt.Sprint(numbers[3]) + fmt.Sprint(numbers[4]) + fmt.Sprint(numbers[5]) + "-" + fmt.Sprint(numbers[6]) + fmt.Sprint(numbers[7]) + fmt.Sprint(numbers[8]) + fmt.Sprint(numbers[9])
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
