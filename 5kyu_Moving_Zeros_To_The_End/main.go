package main

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func MoveZeros(arr []int) []int {
	max := len(arr)
	for i := 0; i < max; i++ {
		if arr[i] == 0 {
			arr = remove(arr, i)
			i--
			max--
			arr = append(arr, 0)
		}
	}

	return arr
}
