package main

import "fmt"

func Bell(n int) (arr []int) {

	arr = append(arr, n)

	if n == 1 {
		return arr
	}
	if n == 2 {
		arr = append(arr, arr[len(arr)-1])
		return arr
	}

	inc := n - 2
	for i := 1; i < int(n/2); i++ {
		arr = append(arr, arr[i-1]+inc)
		inc -= 2
	}
	if n%2 == 1 && n >= 3 {
		arr = append(arr, arr[len(arr)-1]+1)
		for i := len(arr) - 2; i >= 0; i-- {
			arr = append(arr, arr[i])
		}
	} else {
		for i := len(arr) - 1; i >= 0; i-- {
			arr = append(arr, arr[i])
		}
	}

	return arr
}

func main() {
	fmt.Println(Bell(5))
}
