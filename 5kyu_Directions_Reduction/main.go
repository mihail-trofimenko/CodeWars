package main

import "fmt"

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func DirReduc(arr []string) []string {
	opposite := [4][2]string{
		{"NORTH", "SOUTH"},
		{"SOUTH", "NORTH"},
		{"WEST", "EAST"},
		{"EAST", "WEST"}}
	reduced := true

	for reduced {
		reduced = false
		for i := 0; i < len(arr)-1; i++ {
			for d := 0; d < 4; d++ {
				if len(arr)-2 < 0 {
					break
				}
				if arr[i] == opposite[d][0] && arr[i+1] == opposite[d][1] {
					arr = remove(arr, i+1)
					arr = remove(arr, i)
					reduced = true
					break
				}
			}
			if reduced {
				break
			}
		}
	}
	return arr
}

func main() {
	a := []string{"EAST", "EAST", "WEST", "NORTH", "WEST", "EAST", "EAST", "SOUTH", "NORTH", "WEST"}

	fmt.Println(DirReduc(a))
}
